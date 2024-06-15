package routes

import (
	"scoreboard/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Public routes
func SetupPrivateRoutes(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the dashboard"})
	})

	router.POST("/updateProfile", func(c *gin.Context) {
		// Handle profile update
		c.JSON(200, gin.H{"message": "Profile updated"})
	})
}
