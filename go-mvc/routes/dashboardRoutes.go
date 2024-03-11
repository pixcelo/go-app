package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/controllers"
)

func DashboardRoutes(ctrl *controllers.DashboardController, r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/dashboard", ctrl.ShowDashboard)
}
