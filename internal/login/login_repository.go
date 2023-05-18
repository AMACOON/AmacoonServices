package login

import (
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewLoginRepository(db *gorm.DB, logger *logrus.Logger) *LoginRepository {
	return &LoginRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *LoginRepository) Login(loginRequest LoginRequest) (*owner.Owner, error) {
	r.Logger.Info("Repository Login")

	user := &owner.Owner{}
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
