package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		// Handle login
		c.JSON(200, gin.H{"message": "login successful"})
	})
}
