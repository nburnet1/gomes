package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetLine(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "line",
    })
}