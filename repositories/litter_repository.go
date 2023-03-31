package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

type LitterRepository struct {
	DB *gorm.DB
}

func (r *LitterRepository) GetAllLitters() ([]models.Litter, error) {
	var litters []models.Litter
	if err := r.DB.Find(&litters).Error; err != nil {
		return nil, err
	}
	return litters, nil
}

func (r *LitterRepository) GetLitterByID(id uint) (*models.Litter, []*models.Kitten, error) {
	var litter models.Litter
	if err := r.DB.First(&litter, id).Error; err != nil {
		return nil, nil, err
	}

	var kittens []*models.Kitten
	if err := r.DB.Where("id_ninhadas = ?", id).Find(&kittens).Error; err != nil {
		return nil, nil, err
	}

	return &litter, kittens, nil
}

func (r *LitterRepository) CreateLitter(litter *models.Litter, kittens []*models.Kitten) (uint, error) {

	// Define o status da ninhada como "Pending"
	litter.Status = "Pending"

	// Define o status dos filhotes como "Pending"
	for _, kitten := range kittens {
		kitten.Status = "Pending"
	}

	tx := r.DB.Begin()

	// Cria a ninhada
	if err := tx.Create(litter).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	// Define o LitterID
	for i, kitten := range kittens {
		kitten.LitterID = litter.ID
		kittens[i] = kitten
	}

	// Cria os gatos da ninhada
	if err := tx.Create(&kittens).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return litter.ID, nil
}

func (r *LitterRepository) UpdateLitter(litter *models.Litter, kittens []*models.Kitten) error {
	tx := r.DB.Begin()

	// Atualiza a ninhada
	if err := tx.Save(litter).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Atualiza os filhotes
	for _, kitten := range kittens {
		if err := tx.Save(kitten).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *LitterRepository) DeleteLitter(litter *models.Litter) error {
	// Exclui os filhotes associados à ninhada
	if err := r.DB.Where("id_ninhadas = ?", litter.ID).Delete(&models.Kitten{}).Error; err != nil {
		return err
	}

	// Exclui a ninhada
	if err := r.DB.Delete(litter).Error; err != nil {
		return err
	}
	return nil
}

func (r *LitterRepository) GetKittensByLitterID(litterID uint) ([]*models.Kitten, error) {
	var kittens []*models.Kitten
	if err := r.DB.Where("id_ninhadas = ?", litterID).Find(&kittens).Error; err != nil {
		return nil, err
	}
	return kittens, nil
}
