package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// GetEnterprise godoc
// @Summary Get enterprise
// @Description Get enterprise
// @Tags enterprise
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /enterprise [get]
func GetEnterprise(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "enterprise",
    })
}