package main

import (
	_ "scoreboard/config"
	_ "scoreboard/database"
	"scoreboard/internal/routes"
	"scoreboard/logger"
)

// @title cricket scoreboard API's
// @version 1.0
// @description API server for cricket scoreboard
// @contact.name Jaskaran Singh

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
	logger.Log.Fatal("Failed to start server")
}
