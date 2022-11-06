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
	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()

	if err != nil {
		panic("Error loading app.env config file: ")
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
