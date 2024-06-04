package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(version *gin.RouterGroup) {
	version.GET("/enterprise", GetEnterprise)
	version.GET("/area", GetArea)
	version.GET("/site", GetSite)
	version.GET("/line", GetLine)
	version.GET("/ping", PingHandler)
	version.GET("/cell", GetCell)
}
