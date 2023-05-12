package litter

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LitterRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewLitterRepository(db *gorm.DB, logger *logrus.Logger) *LitterRepository {
	return &LitterRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *LitterRepository) CreateLitter(litter Litter) (Litter, error) {
	r.Logger.Infof("Repository CreateLitter")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the Litter record
	if err := tx.Create(&litter).Error; err != nil {
		tx.Rollback()
		return Litter{}, err
	}

	// If everything goes well, commit the transaction
	tx.Commit()

	r.Logger.Infof("Repository CreateLitter OK")
	return litter, nil
}

func (r *LitterRepository) GetLitterByID(id uint) (Litter, error) {
	r.Logger.Infof("Repository GetLitterByID")
	var litter Litter
	
	err := r.DB.Preload("KittenData").First(&litter, id).Error
	if err != nil {
		return Litter{}, err
	}

	r.Logger.Infof("Repository GetLitterByID OK")
	return litter, nil
}

func (r *LitterRepository) UpdateLitter(id uint, litter Litter) error {
    r.Logger.Infof("Repository UpdateLitter")
    litter.ID = id

    if err := r.DB.Save(&litter).Error; err != nil {
        r.Logger.Errorf("error updating litter: %v", err)
        return err
    }

    r.Logger.Infof("Repository UpdateLitter OK")
    return nil
}

func (r *LitterRepository) UpdateLitterStatus(id uint, status string) error {
    r.Logger.Infof("Repository UpdateLitterStatus")
    
    err := r.DB.Model(&Litter{}).Where("id = ?", id).Update("status", status).Error
    if err != nil {
        return err
    }
    
    r.Logger.Infof("Repository UpdateLitterStatus OK")
    return nil
}

func (r *LitterRepository) GetAllLittersByRequesterID(requesterID uint) ([]Litter, error) {
	r.Logger.Infof("Repository GetAllLittersByRequesterID")

	var litters []Litter
	if err := r.DB.Where("requester_id = ?", requesterID).Find(&litters).Error; err != nil {
		return nil, err
	}

	r.Logger.Infof("Repository GetAllLittersByRequesterID OK")
	return litters, nil
}





