package catshowregistration

import (
	

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatShowRegistrationRepository struct {
	DB         *gorm.DB
	Logger     *logrus.Logger
}

func NewCatShowRegistrationRepository(db *gorm.DB, logger *logrus.Logger) *CatShowRegistrationRepository {
	return &CatShowRegistrationRepository{
		DB:         db,
		Logger:     logger,
	}
}

func (r *CatShowRegistrationRepository) CreateCatShowRegistration(registration *Registration) (*Registration, error) {
    r.Logger.Infof("Repository CreateCatShowRegistration")

    // Inicia uma transação
    tx := r.DB.Begin()

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Cria o registro de Registration
    if err := tx.Create(registration).Error; err != nil {
        tx.Rollback()
        r.Logger.WithError(err).Error("Failed to create registration")
        return nil, err
    }

    // Se tudo correr bem, confirma a transação
    tx.Commit()

    r.Logger.Infof("Repository CreateCatShowRegistration OK")
    return registration, nil
}


func (r *CatShowRegistrationRepository) DeleteCatShowRegistrationByID(registrationID uint) error {
    r.Logger.Infof("Repository DeleteCatShowRegistrationByID: %d", registrationID)

    // Inicia uma transação
    tx := r.DB.Begin()

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Realiza a operação de delete
    if err := tx.Where("id = ?", registrationID).Delete(&Registration{}).Error; err != nil {
        tx.Rollback()
        r.Logger.WithError(err).Errorf("Failed to delete registration with ID %d", registrationID)
        return err
    }

    // Se tudo correr bem, confirma a transação
    tx.Commit()

    r.Logger.Infof("Repository DeleteCatShowRegistrationByID %d OK", registrationID)
    return nil
}
