package color

import (
	"gorm.io/gorm"
)

type ColorRepository struct {
	DB *gorm.DB
}

func NewColorRepository(db *gorm.DB) *ColorRepository {
	return &ColorRepository{
		DB: db,
	}
}



func (r *ColorRepository) GetAllColorsByBreed(breedID string) ([]Color, error) {
	var colors []Color
	if err := r.DB.Unscoped().Where("id_raca = ?", breedID).Find(&colors).Error; err != nil {
		return nil, err
	}
	return colors, nil
}
