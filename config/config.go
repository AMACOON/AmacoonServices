package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig(logger *logrus.Logger) {
	// Check which environment we are in
	env := os.Getenv("APP_ENV")

	var configFile string

	switch env {
	case "production":
		logger.Println("Production Config Set")
		configFile = "config.prod"
	case "local":
		logger.Println("Local Config Set")
		configFile = "config.local"
	default:
		logger.Println("Unknown environment, using local config")
		configFile = "config.local"
	}

	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Println("Error reading config file, ", err)
	}

	viper.AutomaticEnv()
}
