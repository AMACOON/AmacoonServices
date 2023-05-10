package country

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)



type CountryRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCountryRepository(db *gorm.DB, logger *logrus.Logger) *CountryRepository {
	return &CountryRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CountryRepository) GetAllCountries() ([]Country, error) {
	r.Logger.Infof("Repository GetAllCountries")

	var countries []Country
	if err := r.DB.Where("is_activated = ?", true).Find(&countries).Error; err != nil {
		return nil, err
	}

	r.Logger.Infof("Repository GetAllCountries OK")
	return countries, nil
}
