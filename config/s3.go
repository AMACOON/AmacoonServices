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
		Region:      aws.String(viper.GetString("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("AWS_ACCESSKEYID"), viper.GetString("AWS_SECRETACCESSKEY"), ""),
	})
	if err != nil {
		logger.WithError(err).Error("Failed to create AWS session")
		return nil, err
	}

	// Crie um novo cliente AWS S3
	svc := s3.New(sess)

	// Teste a conexão listando os buckets
	logger.Info("Testing S3 connection")
	result, err := svc.ListBuckets(nil)
	if err != nil {
		logger.WithError(err).Error("Failed to list S3 buckets")
		return nil, err
	}

	// Log os buckets encontrados
	logger.Info("Successfully connected to S3. Buckets found:")
	for _, bucket := range result.Buckets {
		logger.Infof("Bucket Name: %s", aws.StringValue(bucket.Name))
	}

	return svc, nil
}
