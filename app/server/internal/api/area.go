package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArea godoc
// @Summary Get area
// @Description Get area
// @Tags area
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /area [get]
func GetArea(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "area",
	})
}
