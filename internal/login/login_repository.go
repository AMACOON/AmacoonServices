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

func (r *LoginRepository) Login(loginRequest LoginRequest) (*owner.Owner, bool, error) {
	r.Logger.Info("Repository Login")

	user := &owner.Owner{}
	if err := r.DB.Where("email = ?", loginRequest.Email).First(user).Error; err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect %v", loginRequest.Email)
		return nil, false, err
	}

	// Comparar a senha fornecida com a senha armazenada
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect")
		return nil, false, err
	}

	// A senha está correta, verificar se o usuário está associado ao clubId = 4
	var clubCount int64
	if err := r.DB.Model(&owner.OwnerClub{}).Where("owner_id = ? AND club_id = ?", user.ID, 4).Count(&clubCount).Error; err != nil {
		r.Logger.WithError(err).Errorf("Failed to check for club association")
		return nil, false, err
	}

	isAssociated := clubCount > 0

	// retornar o usuário e se ele está associado
	r.Logger.Info("Repository Login OK")
	return user, isAssociated, nil
}


func (r *LoginRepository) ResetPassword(email string, newPassword string) error {
	r.Logger.Info("Repository ResetPassword")

	// Buscar o usuário pelo email
	user := &owner.Owner{}
	if err := r.DB.Where("email = ?", email).First(user).Error; err != nil {
		r.Logger.WithError(err).Errorf("User not found: %v", email)
		return err
	}

	// Gerar o hash da nova senha
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		r.Logger.WithError(err).Error("Failed to generate password hash")
		return err
	}

	// Atualizar a senha do usuário no banco de dados
	user.PasswordHash = string(newHashedPassword)
	if err := r.DB.Save(user).Error; err != nil {
		r.Logger.WithError(err).Error("Failed to save user with new password")
		return err
	}

	r.Logger.Info("Repository ResetPassword OK")
	return nil
}

