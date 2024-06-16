package routes

import (
	"scoreboard/internal/handlers/auth"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the scoreboard"})
	})

	// Auth routes
	authGroup := router.Group("/auth")

	auth.SetupAuthHandlers(authGroup)

}
