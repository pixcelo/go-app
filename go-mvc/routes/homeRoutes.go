package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/controllers"
)

type HomeController struct{}

func HomeRoutes(ctrl *controllers.HomeController, r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/", ctrl.ShowHome)
}
