package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/controllers"
)

func LoginRoutes(ctrl *controllers.LoginController, r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/login", ctrl.ShowLogin)
}
