package cat

import (
	
	"github.com/sirupsen/logrus"
)


type CatService struct {
	CatRepo CatRepoInterface
	Logger        *logrus.Logger
}

func NewCatService(catRepo CatRepoInterface, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo: catRepo,
        Logger:       logger,
	}
}

func (s *CatService) GetCatsCompleteByID(id string) (*CatComplete, error) {
    cats, err := s.CatRepo.GetCatCompleteByID(id)
    if err != nil {
        s.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex from repo")
        return nil, err
    }
    return cats, nil
}



func (s *CatService) GetCatsByOwnerAndGender(ownerID string, gender string) ([]*CatComplete, error) {
    cats, err := s.CatRepo.GetAllByOwnerAndGender(ownerID, gender)
    if err != nil {
        s.Logger.WithError(err).Error("Failed to get cats by Owner and Gender from repo")
        return nil, err
    }
    return cats, nil
}

func (s *CatService) GetCatCompleteByRegistration(registration string) (*CatComplete, error) {
    cat, err := s.CatRepo.GetCatCompleteByRegistration(registration)
    if err != nil {
        s.Logger.WithError(err).Errorf("Failed to get cat by registration '%s' from repo", registration)
        return nil, err
    }
    return cat, nil
}

