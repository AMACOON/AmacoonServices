package user

import (
	"github.com/sirupsen/logrus"
)

type UserService struct {
	UserRepo *UserRepository
	Logger   *logrus.Logger
}

func NewUserService(userRepo *UserRepository, logger *logrus.Logger) *UserService {
	return &UserService{
		UserRepo: userRepo,
		Logger:   logger,
	}
}

func (s *UserService) Login(loginRequest LoginRequest) (*User, error) {
	s.Logger.Info("Service Login")

	user, err := s.UserRepo.Login(loginRequest)
	if err != nil {
		return nil, err
	}
	
	s.Logger.Info("Service Login OK")
	return user, nil
}
