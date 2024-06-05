package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSite godoc
// @Summary Get site
// @Description Get site
// @Tags site
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /site [get]
func GetSite(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "site",
	})
}
