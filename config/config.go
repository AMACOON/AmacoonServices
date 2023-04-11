package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	MongoDBUsername string
	MongoDBPassword string
	MongoDBHost     string
	MongoDBPort     string
	MongoDBName     string
	ServerPort      string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Falha ao ler o arquivo de configuração")
	}

	return &Config{
		DBUsername:     viper.GetString("db.username"),
		DBPassword:     viper.GetString("db.password"),
		DBHost:         viper.GetString("db.host"),
		DBPort:         viper.GetString("db.port"),
		DBName:         viper.GetString("db.name"),
		MongoDBUsername: viper.GetString("mongodb.username"),
		MongoDBPassword: viper.GetString("mongodb.password"),
		MongoDBHost:     viper.GetString("mongodb.host"),
		MongoDBPort:     viper.GetString("mongodb.port"),
		MongoDBName:     viper.GetString("mongodb.name"),
		ServerPort:      viper.GetString("server.port"),
	}
}
