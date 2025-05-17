package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nburnet1/gomes/internal/config"
)

func init() {
	endpoints := []config.Endpoint{
		{
			Method:   config.GET,
			Url:      "/admin/users",
			Handlers: []gin.HandlerFunc{showAddUserPage},
		},
		{
			Method:   config.POST,
			Url:      "/admin/users",
			Handlers: []gin.HandlerFunc{handleAddUser},
		},
	}
	config.RegisterEndpoints(endpoints...)
}

// Show the add user page
func showAddUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add_user.html", nil)
}

// Handle user creation
func handleAddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")

	// Validate input
	if username == "" || password == "" || role == "" {
		RenderError(c, http.StatusBadRequest, "All fields are required")
		return
	}

	// Create user object
	// user := database.User{Username: username, Role: role}
	// if err := user.SetPassword(password); err != nil {
	// 	RenderError(c, http.StatusInternalServerError, "Failed to hash password")
	// 	return
	// }

	// // Save user to database
	// if err := database.DB.Create(&user).Error; err != nil {
	// 	RenderError(c, http.StatusInternalServerError, "Failed to create user")
	// 	return
	// }

	c.HTML(http.StatusOK, "add-user.html", gin.H{
		"message": "User created successfully",
	})
}