package country

import (
	

	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
    return &CountryRepository{
        DB: db,
    }
}

func (r *CountryRepository) GetAllCountries() ([]Country, error) {
	var countries []Country
	if err := r.DB.Unscoped().Where("visivel = ?", "s").Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}
