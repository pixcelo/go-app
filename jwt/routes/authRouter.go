package routes

import (
	controller "jwt/controllers"
	"github.com/gin-gonic/gin"
)
func AuthRoutes(incoming Routes *gin.Engine) {
	incomingRoutes.POST("user/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}