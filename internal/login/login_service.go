package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LoginService struct {
	LoginRepo *LoginRepository
	Logger    *logrus.Logger
}

func NewLoginService(loginRepo *LoginRepository, logger *logrus.Logger) *LoginService {
	return &LoginService{
		LoginRepo: loginRepo,
		Logger:    logger,
	}
}

func (s *LoginService) Login(loginRequest LoginRequest) (*LoginResponse, error) {
	s.Logger.Info("Service Login")

	user, isAssociated, err := s.LoginRepo.Login(loginRequest)
	if err != nil {
		return nil, err
	}

	// Gerar o token JWT após a autenticação bem-sucedida
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", user.ID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token válido por 24 horas
	})
	secret := viper.GetString("jwt.secret")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		s.Logger.WithError(err).Error("failed to sign token")
		return nil, err
	}

	s.Logger.Info("Service Login OK")
	return &LoginResponse{
		Owner: user,
		Token: tokenString,
		IsAssociate: isAssociated,
	}, nil
}
