package routes

import (
	"github.com/gin-gonic/gin"
)

func HomeRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("views/*/*")
	r.GET("/", showHome)
}

func showHome(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
