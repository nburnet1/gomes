package api

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/internal/grpcclient"
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
			Method: config.GET,
			Url: "sse/admin/node",
			Handlers: []gin.HandlerFunc{subNode},
		},
	}
	config.RegisterEndpoints(endpoints...)
}

func getNamespace(c *gin.Context) {

	
	namespaceClient := grpcclient.NamespaceClient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	topic := c.Query("topic")

	var err error
	var nodes *pb.BrowseNodesResponse
	defer cancel()
	if topic == "" {
		nodes, err = namespaceClient.BrowseRootNodes(ctx, &pb.BrowseRootNodesRequest{})
	} else {
		nodes, err = namespaceClient.GetChildren(ctx, &pb.ReadNodeRequest{Topic: topic})
	}

	if err != nil {
		fmt.Println("Failed to get root nodes:", err)
		return
	}
	if topic == "" {
		c.HTML(200, "namespace.html", nodes)
	} else {
		c.HTML(200, "node.html", nodes)
	}
}


func getNode(c *gin.Context) {



    topic := c.Query("topic")

    if topic == "" {
        c.JSON(400, gin.H{"error": "No topic provided"})
        return
    }

	NamespaceClient := grpcclient.NamespaceClient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

    node, err := NamespaceClient.ReadNode(ctx, &pb.ReadNodeRequest{Topic: topic})
    if err != nil {
        c.JSON(400, gin.H{"error": "Node not found"})
        return
    }

    c.HTML(200, "edit-node.html", node)
}

func subNode(c *gin.Context) {
	// Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Get the topic from query params
	topic := c.Query("topic")

	// Establish gRPC stream
	namespaceClient := grpcclient.NamespaceClient
	ctx, cancel := context.WithCancel(c.Request.Context()) // Use Gin's request context to handle disconnections
	defer cancel()

	serverStream, err := namespaceClient.SubNode(ctx, &pb.SubNodeRequest{Topic: topic})
	if err != nil {
		c.JSON( 500, "Failed to subscribe to node",)
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
			_, writeErr := fmt.Fprintf(c.Writer, "data: %s\n\n", node)
			if writeErr != nil {
				fmt.Println("Error writing to response:", writeErr)
				return
			}

			// Flush to ensure the event is sent immediately
			c.Writer.Flush()
		}
	}
}