package routes

import (
	"scoreboard/internal/middleware"

	_ "scoreboard/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {

	// router := gin.Default()
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.CORSMiddleware())

	// Set up public routes
	publicRoutes := router.Group("/v1/api/public")
	SetupPublicRoutes(publicRoutes)

	// Set up private routes
	privateRoutes := router.Group("/v1/api/private")
	SetupPrivateRoutes(privateRoutes)

	return router
}
