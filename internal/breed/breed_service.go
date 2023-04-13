package breed

import (
	
	"github.com/sirupsen/logrus"
)

type BreedService struct {
	BreedRepo *BreedRepository
	Logger        *logrus.Logger
}

func NewBreedService(breedRepo *BreedRepository, logger *logrus.Logger) *BreedService {
	return &BreedService{
		BreedRepo: breedRepo,
        Logger:       logger,
	}
}

func (s *BreedService) GetAllBreeds() ([]BreedMongo, error) {
    s.Logger.Infof("Service GetAllBreeds")
    breeds, err := s.BreedRepo.GetAllBreeds()
    if err != nil {
        s.Logger.WithError(err).Error("failed to get all breeds")
        return nil, err
    }
    s.Logger.Infof("Service GetAllBreeds OK")
    return breeds, nil
}

func (s *BreedService) GetBreedByID(id string) (*BreedMongo, error) {
	s.Logger.Infof("Service GetBreedByID")
    breed, err := s.BreedRepo.GetBreedByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching breed from repository: %v", err)
		return nil, err
	}
    s.Logger.Infof("Service GetBreedByID OK")
	return breed, nil
}

