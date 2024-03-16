package config

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type FacebookConfig struct {
	Facebook Facebook `yaml:facebook`
}

type Facebook struct {
	ClientId     string `yaml:clientId`
	ClientSecret string `yaml:clientSecret`
	RedirectURL  string `yaml:redirectURL`
}

func InitFacebookConfig() (*oauth2.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error: Reading config.yml file %s \n", err)
	}

	var config FacebookConfig

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}

	var oauthConfig = &oauth2.Config{
		ClientID:     config.Facebook.ClientId,
		ClientSecret: config.Facebook.ClientSecret,
		RedirectURL:  config.Facebook.RedirectURL,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}

	return oauthConfig, err
}
