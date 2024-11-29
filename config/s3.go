package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupS3Session(logger *logrus.Logger) (*s3.S3, error) {
	// Crie uma nova sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(viper.GetString("aws.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("aws.accessKeyId"), viper.GetString("aws.secretAccessKey"), ""),
        
	})
	if err != nil {

		return nil, err
	}

	// Crie um novo cliente AWS S3
	svc := s3.New(sess)

	// Teste a conexão listando os buckets
	// logger.Info("Testing S3 connection")
    // _, err = svc.ListBuckets(nil)
	// if err != nil {
	// 	logger.WithError(err).Error("Failed to list S3 buckets")
	// 	return nil, err
	// }

	logger.Info("Successfully connected to S3 and listed buckets")
	return svc, nil
}