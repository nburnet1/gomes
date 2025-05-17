package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nburnet1/gomes/internal/client"
	"github.com/nburnet1/gomes/internal/config"

	_ "github.com/nburnet1/gomes/internal/web/api"
)

func main() {
	client.AuthGRPC.Connect("localhost:50051")
	client.NamespaceGRPC.Connect("localhost:50055")
	defer client.NamespaceGRPC.Close()
	defer client.AuthGRPC.Close()
	
	r := gin.Default()
	r.Static("/static", "internal/web/static")
	r.LoadHTMLGlob("internal/web/templates/**/*.gohtml")
	config.EnableEndpointsFromEngine(r)

	// Set Trusted Proxies
	if err := r.SetTrustedProxies(nil); err != nil {
		fmt.Println("Error setting trusted proxies:", err)
	}

	// Run the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
