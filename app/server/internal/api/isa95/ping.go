package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// PingHandler handles the /ping route
func PingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}
