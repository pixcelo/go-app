package main

import (
	"github.com/pixcelo/go-mvc/routes"
)

func main() {
	// Load configurations
	// config.LoadConfig()

	// Initialize database connection
	// db.Connect()

	// Setup Gin router
	router := routes.SetupRoutes()
	router.Run("localhost:8530")
}
