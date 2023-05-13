package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserRepository {
	return &UserRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *UserRepository) Login(loginRequest LoginRequest) (*User, error) {
	r.Logger.Info("Repository Login")
	
	user := &User{}
	if err := r.DB.Where("email = ?", loginRequest.Email).First(user).Error; err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect %v", loginRequest.Email)
		return nil, err
	}

	// Comparar a senha fornecida com a senha armazenada
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect")
		return nil, err
	}

	// A senha está correta, retornar o usuário
	r.Logger.Info("Repository Login OK")
	return user, nil
}
