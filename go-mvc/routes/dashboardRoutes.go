package routes

import (
	"github.com/gin-gonic/gin"
)

// func DashboardRoutes(r *gin.Engine) {
// 	dashboardCtrl := &controllers.DashboardController{}

// 	dashboardRoutes := r.Group("/dashboard")
// 	dashboardRoutes.Use(middleware.AuthRequired)
// 	{
// 		dashboardRoutes.GET("/", dashboardCtrl.ShowDashboard)
// 	}
// }

func DashboardRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/dashboard", showDashboard)
}

func showDashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", nil)
}
