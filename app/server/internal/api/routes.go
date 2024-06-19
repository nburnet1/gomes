package api

import (
	"github.com/gin-gonic/gin"
	"gomes/server/internal/api/isa95"
)

func RegisterRoutes(version *gin.RouterGroup) {
	version.GET("/enterprise", api.GetEnterprise)
	version.GET("/area", api.GetArea)
	version.GET("/site", api.GetSite)
	version.GET("/line", api.GetLine)
	version.GET("/ping", api.PingHandler)
	version.GET("/cell", api.GetCell)
}
