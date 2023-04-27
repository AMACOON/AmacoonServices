package cat

import (
	"github.com/sirupsen/logrus"
)

type CatService struct {
	CatRepo CatRepoInterface
	Logger  *logrus.Logger
}

func NewCatService(catRepo CatRepoInterface, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo: catRepo,
		Logger:  logger,
	}
}

// GetCatsCompleteByID returns a CatComplete by its ID
// @param: id: the ID of the cat
func (s *CatService) GetCatsCompleteByID(id string) (*CatComplete, error) {
	s.Logger.Infof("Service GetCatsCompleteByID")
	cats, err := s.CatRepo.GetCatCompleteByID(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Id from repo")
		return nil, err
	}

	s.Logger.Infof("Service GetCatsCompleteByID OK")
	return cats, nil
}

func (s *CatService) GetCatsByOwnerAndGender(ownerID string, gender string) ([]*CatComplete, error) {
	s.Logger.Infof("Service GetCatsByOwnerAndGender")
	cats, err := s.CatRepo.GetAllByOwnerAndGender(ownerID, gender)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner and Gender from repo")
		return nil, err
	}
	
	s.Logger.Infof("Service GetCatsByOwnerAndGender OK")
	return cats, nil
}

func (s *CatService) GetCatCompleteByRegistration(registration string) (*CatComplete, error) {
	s.Logger.Infof("Service GetCatCompleteByRegistration")
	cat, err := s.CatRepo.GetCatCompleteByRegistration(registration)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to get cat by registration '%s' from repo", registration)
		return nil, err
	}
	s.Logger.Infof("Service GetCatCompleteByRegistration OK")
	return cat, nil
}

func (s *CatService) GetAllByOwner(ownerID string) ([]*CatComplete, error) {
	s.Logger.Infof("Service GetAllByOwner")
	cats, err := s.CatRepo.GetAllByOwner(ownerID)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return nil, err
	}
	s.Logger.Infof("Service GetAllByOwner OK")
	return cats, nil
}


func GetFullName(cat *CatComplete) string {
	var prefix, suffix string

	for _, titleCat := range cat.Titles {
		title := titleCat.Title
		if title.Type == "Championship/Premiorship Titles" {
			prefix += title.Code + " "
		} else if title.Type == "Winner Titles" {
			prefix += titleCat.Date.Format("06") + " " + title.Code + " "
		} else if title.Type == "Merit Titles" {
			suffix += " " + title.Code
		}
	}

	return prefix + cat.Name + suffix
}
// WW'Ano se for 2 anos WW'20'21