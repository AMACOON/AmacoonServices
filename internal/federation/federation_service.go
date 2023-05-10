package federation

import (
	
	"github.com/sirupsen/logrus"
)

type FederationService struct {
	FederationRepo *FederationRepository
	Logger        *logrus.Logger
}

func NewCatteryService(federationRepo *FederationRepository, logger *logrus.Logger) *FederationService {
	return &FederationService{
		FederationRepo: federationRepo,
        Logger:       logger,
	}
}

func (s *FederationService) GetAllFederations() ([]Federation, error) {
    s.Logger.Infof("Service GetAllFederations")
    federations, err := s.FederationRepo.GetAllFederations()
    if err != nil {
        s.Logger.WithError(err).Error("failed to get all federations")
        return nil, err
    }
    s.Logger.Infof("Service GetAllFederations OK")
    return federations, nil
}

func (s *FederationService) GetFederationByID(id string) (*Federation, error) {
	s.Logger.Infof("Service GetFederationByID")
    cattery, err := s.FederationRepo.GetFederationByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching cattery from repository: %v", err)
		return nil, err
	}
    s.Logger.Infof("Service GetFederationByID OK")
	return cattery, nil
}

