package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pixcelo/go-mvc/config"
	"golang.org/x/oauth2"
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

func (ctrl LoginController) AuthFacebookLogin(c *gin.Context) {
	oauthConfig, err := config.InitFacebookConfig()
	if err != nil {
		log.Fatal(err)
	}

	// PKCEを使ってCSRF攻撃から守る
	verifier := oauth2.GenerateVerifier()

	// ユーザーを同意ページにリダイレクトして許可を求める
	url := oauthConfig.AuthCodeURL("state", oauth2.S256ChallengeOption(verifier))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ctrl LoginController) AuthFacebookCallback(c *gin.Context) {
	// oauthConfig, err := config.InitFacebookConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func NewLoginController() *LoginController {
	return &LoginController{}
}
