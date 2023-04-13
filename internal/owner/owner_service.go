package owner

import (
	
	"github.com/sirupsen/logrus"
)

type OwnerService struct {
	OwnerRepo *OwnerRepository
	Logger        *logrus.Logger
}

func NewOwnerService(ownerRepo *OwnerRepository, logger *logrus.Logger) *OwnerService {
	return &OwnerService{
		OwnerRepo: ownerRepo,
		Logger:        logger,
	}
}

func (s *OwnerService) GetOwnerByID(id string) (*OwnerMongo, error) {
	s.Logger.Info("Service GetOwnerByID")
	owner, err := s.OwnerRepo.GetOwnerByExhibitorID(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get owner from repository: %v", err)
		return nil, err
	}
	s.Logger.Info("Service GetOwnerByID OK")
	return owner, nil
}

func (s *OwnerService) GetAllOwners() ([]OwnerMongo, error) {
    s.Logger.Infof("Service GetAllOwners")
    owners, err := s.OwnerRepo.GetAllOwners()
    if err != nil {
        s.Logger.WithError(err).Error("failed to get all owners")
        return nil, err
    }
    s.Logger.Infof("Service GetAllOwners OK")
    return owners, nil
}

