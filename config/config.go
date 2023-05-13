package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	ServerPort      string
	S3AwsAccessKeyId string
	S3AwsSecretAccessKey string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("Falha ao ler o arquivo de configuração")
	}

	return &Config{
		DBUsername:     viper.GetString("db.username"),
		DBPassword:     viper.GetString("db.password"),
		DBHost:         viper.GetString("db.host"),
		DBPort:         viper.GetString("db.port"),
		DBName:         viper.GetString("db.name"),
		ServerPort:      viper.GetString("server.port"),
		S3AwsAccessKeyId: viper.GetString("aws.S3AwsAccessKeyId"),
		S3AwsSecretAccessKey: viper.GetString("aws.S3AwsSecretAccessKey"),
	}
}
