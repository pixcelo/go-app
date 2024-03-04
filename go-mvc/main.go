package main

import (
	"github.com/gin-gonic/gin"

	config "github.com/pixcelo/go-mvc/config"
	controllers "github.com/pixcelo/go-mvc/controllers"
	db "github.com/pixcelo/go-mvc/db"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Initialize database connection
	db.Connect()

	// Setup Gin router
	router := gin.Default()

	// Register routes
	controllers.RegisterRoutes(router)

	router.Run("localhost:8530")
}
