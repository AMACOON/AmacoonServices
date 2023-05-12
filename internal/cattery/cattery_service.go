package cattery

import (
	
	"github.com/sirupsen/logrus"
)

type CatteryService struct {
	CatteryRepo *CatteryRepository
	Logger        *logrus.Logger
}

func NewCatteryService(catteryRepo *CatteryRepository, logger *logrus.Logger) *CatteryService {
	return &CatteryService{
		CatteryRepo: catteryRepo,
        Logger:       logger,
	}
}

func (s *CatteryService) GetAllCatteries() ([]Cattery, error) {
    s.Logger.Infof("Service GetAllCatteries")
    catteries, err := s.CatteryRepo.GetAllCatteries()
    if err != nil {
        s.Logger.WithError(err).Error("failed to get all catteries")
        return nil, err
    }
    s.Logger.Infof("Service GetAllCatteries OK")
    return catteries, nil
}

func (s *CatteryService) GetCatteryByID(id string) (*Cattery, error) {
	s.Logger.Infof("Service GetCatteryByID")
    cattery, err := s.CatteryRepo.GetCatteryByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching cattery from repository: %v", err)
		return nil, err
	}
    s.Logger.Infof("Service GetCatteryByID OK")
	return cattery, nil
}

