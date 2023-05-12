package config



import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
	
func SetupS3Session(config *Config) (*s3.S3, error) {
    // Crie uma nova sessão AWS
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"), // Atualize com a região do seu bucket
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyId, config.AwsSecretAccessKey, ""),
    })
    if err != nil {
		
        return nil, err
    }

    // Crie um novo cliente AWS S3
    svc := s3.New(sess)

    return svc, nil
}
