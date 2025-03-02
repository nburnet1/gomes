package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/internal/grpcclient"

	_ "github.com/nburnet1/gomes/internal/api"
)

func main() {
	grpcclient.InitGRPC()	
	grpcclient.InitAuthGRPC()
	
	r := gin.Default()
	r.Static("/static", "internal/web/static")
	r.LoadHTMLGlob("internal/web/templates/*.html")
	config.EnableEndpointsFromEngine(r)

	// Set Trusted Proxies
	if err := r.SetTrustedProxies(nil); err != nil {
		fmt.Println("Error setting trusted proxies:", err)
	}

	// Run the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}

	defer grpcclient.GrpcConn.Close()
	defer grpcclient.AuthGrpcConn.Close()
}



