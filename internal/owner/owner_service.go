package owner

import (
	
	"github.com/sirupsen/logrus"
)

type OwnerService struct {
	ExhibitorRepo *OwnerRepository
	Logger        *logrus.Logger
}

func NewOwnerService(exhibitorRepo *OwnerRepository, logger *logrus.Logger) *OwnerService {
	return &OwnerService{
		ExhibitorRepo: exhibitorRepo,
		Logger:        logger,
	}
}

func (s *OwnerService) GetOwnerByID(ownerID int) (*Owner, error) {
	s.Logger.Info("Getting owner with ID ", ownerID)
	owner, err := s.ExhibitorRepo.GetOwnerByExhibitorID(uint(ownerID))
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get owner from repository")
		return nil, err
	}

	return &owner, nil
}

