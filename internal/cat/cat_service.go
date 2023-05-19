package cat

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatService struct {
	CatRepo        *CatRepository
	CatFileService *CatFileService
	Logger         *logrus.Logger
}

func NewCatService(catRepo *CatRepository, catFileService *CatFileService, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo:        catRepo,
		CatFileService: catFileService,
		Logger:         logger,
	}
}

func (s *CatService) CreateCat(req *Cat, filesWithDesc []utils.FileWithDescription) (*Cat, error) {
	s.Logger.Infof("Service CreateCat")

	req.Validated = false
	cat, err := s.CatRepo.CreateCat(req)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Save the files for this cat
	filesCat, err := s.CatFileService.SaveCatFiles(cat.ID, filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error saving cat files: %v", err)
		return nil, err
	}

	cat.Files = &filesCat

	s.Logger.Infof("Service CreateCat OK")
	return cat, nil
}

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

func (s *CatService) GetCatsByOwner(ownerID string) ([]CatInfo, error) {
	s.Logger.Infof("Service GetCatsByOwner")

	cats, err := s.CatRepo.GetCatsByOwner(ownerID)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return nil, err
	}
	s.Logger.Infof("Service GetCatsByOwner OK")
	return cats, nil
}

func (s *CatService) UpdateNeuteredStatus(catID string, neutered bool) error {
	s.Logger.Infof("Service UpdateNeuteredStatus")

	err := s.CatRepo.UpdateNeuteredStatus(catID, neutered)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return  err
	}
	s.Logger.Infof("Service UpdateNeuteredStatus OK")
	return nil
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
