package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"encoding/json"
	"fmt"
)

func LoadConfigFromYAML(logger *logrus.Logger) {
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
		logger.WithError(err).Fatal("Error reading config file")
	}

	viper.AutomaticEnv() // Override config file values with environment variables if present
}


// RDSConfig representa as informações contidas no segredo AppDb
type RDSConfig struct {
    Username             string `json:"username"`
    Password             string `json:"password"`
    Engine               string `json:"engine"`
    Host                 string `json:"host"`
    Port                 int    `json:"port"`
    DBInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

func LoadSecrets(logger *logrus.Logger) {
    secretNames := []string{"AppAwsAccessKeyId", "AppAwsSecretAccessKey", "AppJwtSecret","AppDbCat",}
    region := "sa-east-1"
	viper.Set("AppAwsRegion", region)


    awsConfig, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(region),
    )
    if err != nil {
        logger.WithError(err).Fatal("Unable to load SDK config")
    }

    svc := secretsmanager.NewFromConfig(awsConfig)

    for _, secretName := range secretNames {
        input := &secretsmanager.GetSecretValueInput{
            SecretId:     aws.String(secretName),
            VersionStage: aws.String("AWSCURRENT"),
        }

        result, err := svc.GetSecretValue(context.TODO(), input)
        if err != nil {
            logger.WithFields(logrus.Fields{"secretName": secretName}).WithError(err).Fatal("Failed to retrieve secret from AWS Secrets Manager")
            continue
        }

		if secretName == "AppDbCat" {
			var dbConfig RDSConfig
			err := json.Unmarshal([]byte(*result.SecretString), &dbConfig)
			if err != nil {
				logger.WithError(err).Fatal("Error unmarshaling RDS config")
			} else {
				viper.Set("db.username", dbConfig.Username)
				viper.Set("db.password", dbConfig.Password)
				viper.Set("db.engine", dbConfig.Engine)
				viper.Set("db.host", dbConfig.Host)
				viper.Set("db.port", dbConfig.Port)
				viper.Set("db.name", dbConfig.DBInstanceIdentifier)
				logger.Info("RDS configuration loaded into Viper")
			}

		
            // Aqui você pode fazer o que precisa com as informações de configuração do DB
            // Por exemplo, armazenar no Viper ou configurar uma conexão de banco de dados
            logger.WithFields(logrus.Fields{
                "host":     dbConfig.Host,
                "username": dbConfig.Username,
                "dbname":   dbConfig.DBInstanceIdentifier,
            }).Info("Successfully retrieved and parsed RDS credentials")
			} else {
				// Deserializar o JSON para extrair apenas o valor, independentemente da chave
				var secretData map[string]string
				err = json.Unmarshal([]byte(*result.SecretString), &secretData)
				if err != nil {
					logger.WithError(err).Fatal(fmt.Sprintf("Failed to unmarshal secret string for secret: %s", secretName))
					continue
				}
	
				// Como esperamos apenas um par de chave-valor, pegamos o valor diretamente
				for _, value := range secretData {
					viper.Set(secretName, value)
					break // Sair após o primeiro valor, já que estamos assumindo apenas um par de chave-valor
				}
	
				logger.WithFields(logrus.Fields{"secretName": secretName}).Info("Successfully retrieved secret from AWS Secrets Manager")
			}
		}
	}

func LoadConfig(logger *logrus.Logger) {
	LoadConfigFromYAML(logger)
	LoadSecrets(logger)

}

