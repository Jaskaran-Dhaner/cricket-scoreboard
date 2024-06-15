package routes

import (
	"accrualEngine/config"
	"accrualEngine/internal/middleware"

	_ "accrualEngine/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {

	// router := gin.Default()
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.CORSMiddleware())

	router.Run(config.Server.Host + ":" + config.Server.Port)

	// Set up public routes
	publicRoutes := router.Group("/public")
	SetupPublicRoutes(publicRoutes)

	// Set up private routes
	privateRoutes := router.Group("/private")
	SetupPrivateRoutes(privateRoutes)

	return router
}
