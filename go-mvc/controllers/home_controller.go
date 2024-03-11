package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (ctrl *HomeController) ShowHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func NewHomeController() *HomeController {
	return &HomeController{}
}
