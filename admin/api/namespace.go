package api

import (
	"fmt"
	"gomes/admin/model"
	"gomes/namespace"
	"gomes/util"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	endpoints := []namespace.Endpoint{
		{
			Method:   namespace.GET,
			Url:      "/admin/namespace",
			Handlers: []gin.HandlerFunc{getNamespace},
		},
		{
			Method:   namespace.GET,
			Url:      "/admin/node",
			Handlers: []gin.HandlerFunc{getNode},
		},
		{
			Method:   namespace.PUT,
			Url:      "/admin/node",
			Handlers: []gin.HandlerFunc{updateNode},
		},
		{
			Method:   namespace.GET,
			Url:      "/admin/asset/test",
			Handlers: []gin.HandlerFunc{testAsset},
		},
		{
			Method:   namespace.GET,
			Url:      "/admin/time",
			Handlers: []gin.HandlerFunc{getTime},
		},
	}
	namespace.RegisterEndpoints(endpoints...)
}

func getNamespace(c *gin.Context) {
	topic := c.Query("topic")

	nodes := namespace.Engine.BrowseRootNodes()

	if topic != "" {
		node, err := namespace.Engine.ReadNode(topic)
		if err != nil {
			c.JSON(400, gin.H{"error": "Node not found"})
		}
		nodes = node.GetChildren()
	}

	type child struct {
		Topic string
		HasChildren bool
	}

	childMap := make(map[string]child)

	for name, node := range nodes {
		topic := node.GetTopic()
		childMap[name] = child{
			Topic: topic,
			HasChildren: len(node.GetChildren()) > 0,
		}
	}

	if topic == "" {
		c.HTML(200, "namespace.html", childMap)
	} else {
		c.HTML(200, "node.html", childMap)
	}
}

func getNode(c *gin.Context) {
    topic := c.Query("topic")

    if topic == "" {
        c.JSON(400, gin.H{"error": "No topic provided"})
        return
    }

    node, err := namespace.Engine.ReadNode(topic)
    if err != nil {
        c.JSON(400, gin.H{"error": "Node not found"})
        return
    }
	type nodeProp struct {
		Topic string
		Data interface{}
	}
    value := node.GetValue()

	jsonValue, err := util.UnmarshalFromInterface(value)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse JSON"})
		return
	}

	prop := nodeProp{
		Topic: topic,
		Data: jsonValue,
	}

    c.HTML(200, "edit-node.html", prop)
}

func updateNode(c *gin.Context){
	c.Request.ParseForm()
	fmt.Println(c.Request.Form)
	topic := c.Request.Form.Get("topic")
	value := c.Request.Form.Get("String")

	namespace.Engine.UpdateNode(topic, value)

}

func testAsset(c *gin.Context) {
	node, _ := namespace.Engine.ReadNode("Enterprise/Site/Area/Machine1/Asset")
	asset := node.GetValue().(model.Asset)

	asset.OEEEnabled = !asset.OEEEnabled

	namespace.Engine.UpdateNode("Enterprise/Site/Area/Machine1/Asset", asset)
	c.JSON(200, asset)
}

func getTime(c *gin.Context) {
	// Set the correct headers for SSE
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Stream events every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request.Context().Done():
			// Client disconnected, stop sending events
			fmt.Println("Client disconnected")
			return
		case t := <-ticker.C:
			// Send a new event
			msg := fmt.Sprintf("event: timeUpdate\ndata: Current time is %s\n\n", t.Format(time.RFC1123))
			if _, err := c.Writer.Write([]byte(msg)); err != nil {
				fmt.Println("Failed to write SSE:", err)
				return
			}
			c.Writer.Flush() // Make sure the event is sent immediately
		}
	}
}
