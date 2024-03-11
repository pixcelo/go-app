package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	DB *gorm.DB
}

type SocialProvider struct {
	Name string
	URL  string
}

var providers = []SocialProvider{
	{Name: "Google", URL: "auth/google"},
	{Name: "Facebook", URL: "auth/facebook"},
}

func (ctrl LoginController) ShowLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{"providers": providers})
}

func NewLoginController() *LoginController {
	return &LoginController{}
}
