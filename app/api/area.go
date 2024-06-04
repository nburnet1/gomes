package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetArea(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "area",
    })
}