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
        Region: aws.String("us-east-1"), // Atualize com a região do seu bucket
		Credentials: credentials.NewStaticCredentials(viper.GetString("aws.S3AwsAccessKeyId"), viper.GetString("aws.S3AwsSecretAccessKey"), ""),
    })
    if err != nil {
		
        return nil, err
    }

    // Crie um novo cliente AWS S3
    svc := s3.New(sess)

    return svc, nil
}

