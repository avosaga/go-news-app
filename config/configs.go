package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	ApyKey string
}

var config *Config

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic("Error loading .env config file: ")
	}
}

func GetConfig() *Config {
	if config == nil {
		port := viper.Get("PORT")

		if port == nil {
			port = "3000"
		}

		config = &Config{
			":" + fmt.Sprintf("%v", port),
			fmt.Sprintf("%s", viper.Get("API_KEY")),
		}
	}

	return config
}
