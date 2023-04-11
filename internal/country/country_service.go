package country

import (
	
	"github.com/sirupsen/logrus"
)

type CountryService struct {
	CountryRepo *CountryRepository
	Logger        *logrus.Logger
}

func NewCountryService(countryRepo *CountryRepository, logger *logrus.Logger) *CountryService {
	return &CountryService{
		CountryRepo: countryRepo,
		Logger:       logger,
	}
}

func (s *CountryService) GetAllCountries() ([]Country, error) {
	s.Logger.Infof("GetAllCountries Sercice")
	countries, err := s.CountryRepo.GetAllCountries()
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get all countries")
		return nil, err
	}
	s.Logger.Infof("GetAllCountries Sercice OK")
	return countries, nil
}



