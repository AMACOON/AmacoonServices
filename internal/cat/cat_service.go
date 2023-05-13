package cat

import (
	"github.com/sirupsen/logrus"
)

type CatService struct {
	CatRepo *CatRepository
	Logger  *logrus.Logger
}

func NewCatService(catRepo *CatRepository, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo: catRepo,
		Logger:  logger,
	}
}

// GetCatsCompleteByID returns a CatComplete by its ID
// @param: id: the ID of the cat
func (s *CatService) GetCatsCompleteByID(id string) (*Cat, error) {
	s.Logger.Infof("Service GetCatsCompleteByID")
	cats, err := s.CatRepo.GetCatCompleteByID(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Id from repo")
		return nil, err
	}

	s.Logger.Infof("Service GetCatsCompleteByID OK")
	return cats, nil
}



func (s *CatService) GetCatCompleteByAllByOwner(ownerID string) ([]*Cat, error) {
	s.Logger.Infof("Service GetCatCompleteByAllByOwner")
	cats, err := s.CatRepo.GetCatCompleteByAllByOwner(ownerID)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return nil, err
	}
	s.Logger.Infof("Service GetCatCompleteByAllByOwner OK")
	return cats, nil
}


// func GetFullName(cat *CatComplete) string {
// 	var prefix, suffix string

// 	for _, titleCat := range cat.Titles {
// 		title := titleCat.Title
// 		if title.Type == "Championship/Premiorship Titles" {
// 			prefix += title.Code + " "
// 		} else if title.Type == "Winner Titles" {
// 			prefix += titleCat.Date.Format("06") + " " + title.Code + " "
// 		} else if title.Type == "Merit Titles" {
// 			suffix += " " + title.Code
// 		}
// 	}

// 	return prefix + cat.Name + suffix
// }
// // WW'Ano se for 2 anos WW'20'21