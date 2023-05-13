package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/owner"
)

func MigrateOwnersToUsers(db *gorm.DB, logger *logrus.Logger) error {
	logger.Infof("Migrating owners to users...")

	var owners []owner.Owner
	if err := db.Find(&owners).Error; err != nil {
		logger.WithError(err).Error("Failed to get owners from database")
		return err
	}

	for _, owner := range owners {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(owner.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			logger.WithError(err).Errorf("Failed to hash password for owner %s", owner.Email)
			return err
		}

		user := User{
			OwnerID:      &owner.ID,
			Email:        owner.Email,
			PasswordHash: string(hashedPassword),
			Name:         owner.Name,
			CPF:          owner.CPF,
			IsAdmin:      false, // Inicialmente, todos os usuários não são administradores
		}

		var count int64
		db.Model(&User{}).Where("email = ?", user.Email).Count(&count)
		if count == 0 {
			if err := db.Create(&user).Error; err != nil {
				logger.WithError(err).Errorf("Failed to migrate owner %s", owner.Email)
				return err
			}
			logger.Infof("Owner %s migrated to user", owner.Email)
		} else {
			logger.Infof("User %s already exists in users table", user.Email)
		}
	}

	logger.Infof("Owners migration to users completed successfully")
	return nil
}


func RemovePasswordHashFromOwners(db *gorm.DB, logger *logrus.Logger) error {
	logger.Infof("Removing PasswordHash from owners...")

	if err := db.Migrator().DropColumn(&owner.Owner{}, "PasswordHash"); err != nil {
		logger.WithError(err).Error("Failed to remove PasswordHash from owners")
		return err
	}

	logger.Infof("PasswordHash removed from owners successfully")
	return nil
}

