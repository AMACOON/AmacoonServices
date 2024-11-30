package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type SmtpService struct {
	Logger *logrus.Logger
}

func NewSmtpService(logger *logrus.Logger) *SmtpService {
	return &SmtpService{
		Logger: logger,
	}
}

func (s *SmtpService) SendEmail(from string, to []string, subject string, body string) error {
	s.Logger.Infof("Sending email from %s to %s", from, to)
	
	// Configuração do servidor SMTP
	smtpHost := viper.GetString("SMTP_HOST")
	smtpPort := viper.GetInt("SMTP_PORT")
	smtpUsername := viper.GetString("SMTP_USERNAME")
	smtpPassword := viper.GetString("SMTP_PASSWORD")

	// Cria uma nova mensagem usando a biblioteca GoMail
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	// Cria um novo objeto Dialer para estabelecer a conexão SMTP
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Envia o e-mail usando o Dialernao, faz sentigo
	if err := dialer.DialAndSend(message); err != nil {
		s.Logger.Errorf("Fail to send e-mail: %v", err)
		return err
	}

	s.Logger.Println("Email sent successfully")
	return nil
}
