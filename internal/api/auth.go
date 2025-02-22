package api

import (
	"context"
	"net/http"


	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/internal/grpcclient"
	pb "github.com/nburnet1/gomes/proto"

	"github.com/gin-gonic/gin"
)

func init() {
	endpoints := []config.Endpoint{
		{
			Method:   config.GET,
			Url:      "/login",
			Handlers: []gin.HandlerFunc{showLoginPage},
		},
		{
			Method:   config.POST,
			Url:      "/login",
			Handlers: []gin.HandlerFunc{handleLogin},
		},
		{
			Method:   config.GET,
			Url:      "/dashboard",
			Handlers: []gin.HandlerFunc{authenticateMiddleware(), showDashboard},
		},
	}
	config.RegisterEndpoints(endpoints...)
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	authClient := grpcclient.AuthClient

	// Call gRPC Login
	resp, err := authClient.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Store token in a cookie
	c.SetCookie("jwt", resp.Token, 3600, "/", "", false, true)
	c.HTML(http.StatusOK, "dashboard.html", resp)
}

func authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		authClient := grpcclient.AuthClient

		// Validate token via gRPC
		resp, err := authClient.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token})
		if err != nil || !resp.Valid {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Set("username", resp.Username)
		c.Set("role", resp.Role)
		c.Next()
	}
}

func showDashboard(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"username": username,
		"role":     role,
	})
}