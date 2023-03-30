package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

// Tabela Gatos tem id_expositor, q eh o proprietario, como eu chego nos dados do proprietario ? esse mesmo id ?

type OwnerRepository struct {
	DB *gorm.DB
}

func (r *OwnerRepository) GetOwnerByExhibitorID(idExhibitor uint) (models.Owner, error) {
	var owner models.Owner
	if err := r.DB.Unscoped().Where("id_expositores = ?", idExhibitor).Find(&owner).Error; err != nil {
		return owner, err
	}
	return owner, nil
}
