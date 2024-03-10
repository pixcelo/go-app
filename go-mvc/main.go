package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/routes"
)

func main() {
	// Load configurations
	// config.LoadConfig()

	// Initialize database connection
	// db.Connect()

	// Setup Gin router
	router := gin.Default()
	routes.HomeRoutes(router)
	routes.LoginRoutes(router)

	router.Run("localhost:8530")
}
