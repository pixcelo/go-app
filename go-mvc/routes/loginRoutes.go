package routes

import (
	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/login", showLogin)
}

func showLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
