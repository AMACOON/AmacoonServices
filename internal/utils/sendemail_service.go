package utils

import (
	"github.com/scuba13/AmacoonServices/config"
	"github.com/sirupsen/logrus"
	 "gopkg.in/gomail.v2"
	"strconv"
)

type SmtpService struct {
	Logger *logrus.Logger
	Config *config.Config
}

func NewSmtpService(config *config.Config, logger *logrus.Logger) *SmtpService {
	return &SmtpService{
		Config: config,
		Logger: logger,
	}
}

func (s *SmtpService) SendEmail(from string, to []string, subject string, body string) error {
	// Configuração do servidor SMTP
	smtpHost := s.Config.SMPTHost
	smtpPort, _ := strconv.Atoi(s.Config.SMTPPort)
	smtpUsername := s.Config.SMTPUsername
	smtpPassword := s.Config.SMTPPassword

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
		s.Logger.Errorf("Falha ao enviar o e-mail: %v", err)
		return err
	}

	s.Logger.Println("E-mail enviado com sucesso")
	return nil
}
