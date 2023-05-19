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

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Atualizar os campos específicos da ninhada na tabela "service_litters"
	if err := tx.Save(&litter).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error updating litter with id %v", id)
		return err
	}
	

	// Atualizar os campos específicos dos gatinhos na tabela "service_kittens_litters"
	for _, updatedKitten := range *litter.KittenData {
		if err := tx.Model(&KittenLitter{}).Where("id = ?", updatedKitten.ID).Updates(updatedKitten).Error; err != nil {
			tx.Rollback()
			r.Logger.WithError(err).Errorf("error updating kitten litter record with id %v", updatedKitten.ID)
			return err
		}
	}

	// If everything goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.WithError(err).Errorf("error committing transaction")
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

func (r *LitterRepository) DeleteLitter(id uint) error {
	r.Logger.Infof("Repository DeleteLitter")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var litter Litter
	if err := tx.First(&litter, id).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error finding litter with id %v", id)
		return err
	}

	// Delete the kittens associated with this litter from "service_kittens_litters"
	if err := tx.Where("litter_id = ?", id).Delete(&KittenLitter{}).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting kitten litter records with litter id %v", id)
		return err
	}

	// Delete the litter from "service_litters"
	if err := tx.Delete(&litter).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting litter with id %v", id)
		return err
	}

	// If everything goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.WithError(err).Errorf("error committing transaction")
		return err
	}

	r.Logger.Infof("Repository DeleteLitter OK")
	return nil
}





