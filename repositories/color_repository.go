package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

type ColorRepository struct {
	DB *gorm.DB
}



func (r *ColorRepository) GetAllColorsByBreed(breedID string) ([]models.Color, error) {
    var colors []models.Color
    if err := r.DB.Unscoped().Where("id_raca = ?", breedID).Find(&colors).Error; err != nil {
        return nil, err
    }
    return colors, nil
}