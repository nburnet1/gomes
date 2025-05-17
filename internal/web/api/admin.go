package api

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/internal/client"
	"github.com/nburnet1/gomes/internal/util"
	pb "github.com/nburnet1/gomes/proto"

	"github.com/gin-gonic/gin"
)

func init() {

	endpoints := []config.Endpoint{
		{
			Method:   config.GET,
			Url:      "/admin/namespace",
			Handlers: []gin.HandlerFunc{getNamespace},
		},
		{
			Method:   config.GET,
			Url:      "/admin/node",
			Handlers: []gin.HandlerFunc{getNode},
		},
		{
			Method:   config.GET,
			Url:      "sse/admin/node",
			Handlers: []gin.HandlerFunc{subNode},
		},
		{
			Method:   config.GET,
			Url:      "/admin/node/stash",
			Handlers: []gin.HandlerFunc{stashNode},
		},
		{
			Method:   config.GET,
			Url:      "/test",
			Handlers: []gin.HandlerFunc{func(c *gin.Context) { c.HTML(200, "admin/test", nil) }},
		},
		
	}
	config.RegisterEndpoints(endpoints...)
}

func getNamespace(c *gin.Context) {
	namespaceClient := client.NamespaceGRPC.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	topic := c.Query("topic")

	var (
		nodes *pb.BrowseNodesResponse
		err   error
	)

	if topic == "" {
		nodes, err = namespaceClient.BrowseRootNodes(ctx, &pb.BrowseRootNodesRequest{})
	} else {
		nodes, err = namespaceClient.GetChildren(ctx, &pb.ReadNodeRequest{Topic: topic})
	}

	if err != nil {
		fmt.Println("Failed to get nodes:", err)
		return
	}

	jsonNodes := make([]JsonNode, len(nodes.Nodes))
	for i, node := range nodes.Nodes {
		jsonNodes[i] = JsonNode{
			Topic:       node.Topic,
			Name:        node.Name,
			ParentTopic: node.ParentTopic,
			Timestamp:   node.Timestamp,
			Value:       string(node.Value),
		}
	}

	// **Sort nodes: `_folder` items first, then alphabetically**
	slices.SortStableFunc(jsonNodes, func(a, b JsonNode) int {
		isFolderA := a.Value == "\"_folder\""
		isFolderB := b.Value == "\"_folder\""

		switch {
		case isFolderA && !isFolderB:
			return -1 // `_folder` items go first
		case !isFolderA && isFolderB:
			return 1 // Non-folder items go after `_folder`
		default:
			return strings.Compare(a.Name, b.Name) // Alphabetical order
		}
	})

	// Determine template
	template := "admin/namespace"
	if topic != "" {
		template = "admin/node"
	}
	c.HTML(200, template, jsonNodes)
}

func getNode(c *gin.Context) {

	topic := c.Query("topic")

	if topic == "" {
		RenderError(c, 400, "No topic provided")
		return
	}

	NamespaceClient := client.NamespaceGRPC.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	node, err := NamespaceClient.ReadNode(ctx, &pb.ReadNodeRequest{Topic: topic})
	if err != nil {
		RenderError(c, 400, "Node not found")
		return
	}
	value, err := util.BytesToJson(node.Value)
	if err != nil {
		RenderError(c, 500, "Failed to convert value to JSON")
		return
	}

	jsonNode := JsonNode{
		Topic:       node.Topic,
		Name:        node.Name,
		ParentTopic: node.ParentTopic,
		Timestamp:   node.Timestamp,
		Value:       string(node.Value),
		Type:        fmt.Sprintf("%T", value),
	}

	c.HTML(200, "admin/edit-node", jsonNode)
}

func stashNode(c *gin.Context) {
	topic := c.Query("topic")

	if topic == "" {
		RenderError(c, 400, "No topic provided")
		return
	}
	NamespaceClient := client.NamespaceGRPC.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	node, err := NamespaceClient.ReadNode(ctx, &pb.ReadNodeRequest{Topic: topic})
	if err != nil {
		RenderError(c, 400, "Node not found")
		return
	}
	value, err := util.BytesToJson(node.Value)
	if err != nil {
		RenderError(c, 500, "Failed to convert value to JSON")
		return
	}

	jsonNode := JsonNode{
		Topic:       node.Topic,
		Name:        node.Name,
		ParentTopic: node.ParentTopic,
		Timestamp:   node.Timestamp,
		Value:       string(node.Value),
		Type:        fmt.Sprintf("%T", value),
	}

	c.HTML(200, "admin/stash", jsonNode)

}

func subNode(c *gin.Context) {
	// Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Create a channel for SSE messages
	messageChan := make(chan string)

	// Get the topic from query params
	topic := c.Query("topic")
	fmt.Println("Subscribing to topic:", topic)

	// Establish gRPC stream
	namespaceClient := client.NamespaceGRPC.Client
	ctx, cancel := context.WithCancel(c.Request.Context()) // Use Gin's request context to handle disconnections
	defer cancel()

	// Fetch all available nodes
	allNodesRequest, err := namespaceClient.BrowseNodes(ctx, &pb.BrowseNodesRequest{})
	if err != nil {
		fmt.Println("Failed to get all nodes:", err)
		return
	}

	// Start a goroutine to write SSE messages
	go func() {
		for msg := range messageChan {
			_, writeErr := fmt.Fprintf(c.Writer, "%s\n", msg)
			if writeErr != nil {
				fmt.Println("Error writing SSE event:", writeErr)
				break
			}
			c.Writer.Flush()
		}
	}()

	// Iterate over all nodes and subscribe to them
	for _, node := range allNodesRequest.Nodes {
		go func(nodeTopic string) {
			fmt.Println("Subscribing to:", nodeTopic)

			serverStream, err := namespaceClient.SubNode(ctx, &pb.SubNodeRequest{Topic: nodeTopic})
			if err != nil {
				messageChan <- fmt.Sprintf("event: error\ndata: Failed to subscribe to %s\n\n", nodeTopic)
				return
			}

			// Stream messages from the node
			for {
				select {
				case <-ctx.Done(): // Stop if client disconnects
					fmt.Println("Client disconnected from", nodeTopic)
					return
				default:
					node, err := serverStream.Recv()
					if err != nil {
						fmt.Println("Error receiving from stream:", err)
						return
					}

					// Send SSE event to the message channel
					messageChan <- fmt.Sprintf("event: updated-%s\ndata: %s\n\n", node.Topic, node.Value)
				}
			}
		}(node.Topic)
	}

	// Wait for client disconnect
	<-c.Writer.CloseNotify()
	close(messageChan) // Cleanup when client disconnects
}
