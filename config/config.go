package config

import (
	
	"strings"

	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

func LoadConfig(logger *logrus.Logger) {
	// Configura o Viper para carregar variáveis do arquivo .env
	viper.SetConfigName(".env")   // Nome do arquivo
	viper.SetConfigType("env")    // Tipo do arquivo
	viper.AddConfigPath(".")      // Caminho onde o arquivo está localizado (raiz do projeto)

	// Tenta ler o arquivo .env
	err := viper.ReadInConfig()
	if err != nil {
		logger.Println("No .env file found. Using system environment variables.")
	}

	// Configura o Viper para usar variáveis de ambiente
	viper.AutomaticEnv()

	// Substitui "." por "_" para compatibilidade com variáveis de ambiente
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	logger.Println("Configuration loaded")
}
