package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetCell(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "cell",
    })
}