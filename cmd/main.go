package main

import (
	_ "accrualEngine/config"
	_ "accrualEngine/database"
	"accrualEngine/internal/routes"
	"accrualEngine/logger"
)

// @title accrual engine API
// @version 1.0
// @description API server for accrual engine
// @contact.name Alexey

// @host localhost:8080
// @BasePath /v1/api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	// Init logger
	if logger.Log == nil {
		panic("Failed to initialize logger")
	}
	logger.Log.Info("accrual engine started")
}

func main() {
	// Run server
	routes.SetupRoutes().Run()
	logger.Log.Fatal("accrual engine stopped")
}
