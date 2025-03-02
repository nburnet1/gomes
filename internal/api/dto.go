package api

import (
	"github.com/gin-gonic/gin"
)

type JsonNode struct {
	Topic string `json:"topic"`
	Value any `json:"value"`
	ParentTopic string `json:"parentTopic"`
	Name string `json:"name"`
	Timestamp string `json:"timestamp"`
	Type string `json:"type"`
}

// RenderError renders the error.html template with the given error message
func RenderError(c *gin.Context, statusCode int, errorMessage string) {
	c.HTML(statusCode, "error.html", gin.H{
		"ErrorMessage": errorMessage,
	})
}