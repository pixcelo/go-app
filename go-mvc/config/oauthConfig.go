package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func InitFacebookConfig() *oauth2.Config {
	var oauthConfig = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}

	return oauthConfig
}
