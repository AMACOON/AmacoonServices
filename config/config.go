package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig carrega as configurações do arquivo .env e das variáveis de ambiente
func LoadConfig(logger *logrus.Logger) (map[string]interface{}, error) {
	// Configura o Viper para carregar variáveis do arquivo .env
	viper.SetConfigFile(".env")   // Define o caminho completo do arquivo .env
	viper.SetConfigType("env")    // Tipo do arquivo

	// Tenta ler o arquivo .env
	err := viper.ReadInConfig()
	if err != nil {
		logger.Warn("Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	} else {
		logger.Info("Arquivo .env carregado com sucesso.")
	}

	// Configura o Viper para usar variáveis de ambiente
	viper.AutomaticEnv()

	// Substitui "." por "_" para compatibilidade com variáveis de ambiente
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Define um prefixo opcional para as variáveis de ambiente
	// Por exemplo, APP_PORT, APP_DATABASE_URL, etc.
	// Se não desejar usar prefixo, pode comentar a linha abaixo
	// viper.SetEnvPrefix("APP")

	// Define valores padrão (opcional)
	viper.SetDefault("PORT", "8080")

	// Obtém todas as configurações carregadas
	configMap := viper.AllSettings()

	logger.Info("Configuração carregada com sucesso.")
	return configMap, nil
}
