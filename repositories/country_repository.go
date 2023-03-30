package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func (r *CountryRepository) GetAllCountries() ([]models.Country, error) {
	var countries []models.Country
	if err := r.DB.Unscoped().Where("visivel = ?", "s").Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}
