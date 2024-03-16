package main

import (
	"github.com/pixcelo/go-mvc/config"
	"github.com/pixcelo/go-mvc/routes"
)

func main() {
	// Initialize database connection
	config.InitDB()

	// Setup Gin router
	router := routes.SetupRoutes()
	router.Run("localhost:8080")
}
