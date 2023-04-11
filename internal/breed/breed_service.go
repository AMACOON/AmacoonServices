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


func (s *BreedService) GetAllBreeds() ([]Breed, error) {
    breeds, err := s.BreedRepo.GetAllBreeds()
    if err != nil {
        s.Logger.WithError(err).Error("Failed to get all breeds")
        return nil, err
    }
    return breeds, nil
}

func (s *BreedService) GetCompatibleBreeds(breedID string) ([]string, error) {
    compatibleBreeds, err := s.BreedRepo.GetCompatibleBreeds(breedID)
    if err != nil {
        s.Logger.WithError(err).Errorf("Failed to get compatible breeds for breed with ID %s", breedID)
        return nil, err
    }
    return compatibleBreeds, nil
}

