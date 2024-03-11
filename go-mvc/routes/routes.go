package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/controllers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Controllers
	homeCtrl := controllers.NewHomeController()
	loginCtrl := controllers.NewLoginController()
	// dashboardCtrl := controllers.NewDashboardController()

	// Routing
	HomeRoutes(homeCtrl, router)
	LoginRoutes(loginCtrl, router)
	// DashboardRoutes(dashboardCtrl, router)

	return router
}
