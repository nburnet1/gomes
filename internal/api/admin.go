package api

import (
	"context"
	"fmt"
	"io"
	"slices"
	"strings"
	"time"

	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/internal/grpcclient"
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
		
	}
	config.RegisterEndpoints(endpoints...)
}

func getNamespace(c *gin.Context) {
	namespaceClient := grpcclient.NamespaceClient
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
	template := "namespace.html"
	if topic != "" {
		template = "node.html"
	}
	c.HTML(200, template, jsonNodes)
}

func getNode(c *gin.Context) {

	topic := c.Query("topic")

	if topic == "" {
		RenderError(c, 400, "No topic provided")
		return
	}

	NamespaceClient := grpcclient.NamespaceClient
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

	c.HTML(200, "edit-node.html", jsonNode)
}

func stashNode(c *gin.Context) {
	topic := c.Query("topic")

	if topic == "" {
		RenderError(c, 400, "No topic provided")
		return
	}

	NamespaceClient := grpcclient.NamespaceClient
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

	c.HTML(200, "stash.html", jsonNode)

}

func subNode(c *gin.Context) {
	// Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Get the topic from query params
	topic := c.Query("topic")
	fmt.Println("Subscribing to topic:", topic)
	// Establish gRPC stream
	namespaceClient := grpcclient.NamespaceClient
	ctx, cancel := context.WithCancel(c.Request.Context()) // Use Gin's request context to handle disconnections
	defer cancel()

	allNodesRequest, err := namespaceClient.BrowseNodes(ctx, &pb.BrowseNodesRequest{})
	if err != nil {
		fmt.Println("Failed to get all nodes:", err)
		return
	}

	for _, node := range allNodesRequest.Nodes {
		fmt.Println("Node:", node.Topic)
	}

	
	serverStream, err := namespaceClient.SubNode(ctx, &pb.SubNodeRequest{Topic: topic})
	if err != nil {
		// For SSE, we need to send the error as an SSE event
		_, writeErr := fmt.Fprintf(c.Writer, "event: error\ndata: Failed to subscribe to node\n\n")
		if writeErr != nil {
			fmt.Println("Error writing to response:", writeErr)
		}
		c.Writer.Flush()
		return
	}

	// Create a channel to detect when the client disconnects
	notify := c.Writer.CloseNotify()

	for {
		select {
		case <-notify:
			fmt.Println("Client disconnected")
			return
		default:
			// Read messages from gRPC stream
			node, err := serverStream.Recv()

			if err != nil {
				if err == io.EOF {
					fmt.Println("Stream closed by server")
				} else {
					fmt.Println("Error receiving from stream:", err)
				}
				return
			}

			// Send SSE event to client
			_, writeErr := fmt.Fprintf(c.Writer, "event: updated-%s\ndata: %s\n\n", node.Topic,node.Value)

			if writeErr != nil {
				fmt.Println("Error writing to response:", writeErr)
				return
			}

			// Flush to ensure the event is sent immediately
			c.Writer.Flush()
		}
	}
}
