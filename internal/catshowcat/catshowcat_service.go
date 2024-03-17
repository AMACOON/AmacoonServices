package catshowcat

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatShowCatService struct {
	CatShowCatRepo        *CatShowCatRepository
	CatFileService *CatShowCatFileService
	Logger         *logrus.Logger
}

func NewCCatShowtService(catShowcatRepo *CatShowCatRepository, catFileService *CatShowCatFileService, logger *logrus.Logger) *CatShowCatService {
	return &CatShowCatService{
		CatShowCatRepo:        catShowcatRepo,
		CatFileService: catFileService,
		Logger:         logger,
	}
}

func (s *CatShowCatService) CreateCatShowCat(req *CatShowCat, filesWithDesc []utils.FileWithDescription) (*CatShowCat, error) {
	s.Logger.Infof("Service CreateCatShowCat")

	req.Validated = false
	cat, err := s.CatShowCatRepo.CreateCatShowCat(req)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Check if there are files to save
	if len(filesWithDesc) > 0 {
		// Save the files for this cat
		filesCat, err := s.CatFileService.SaveCatShowCatFiles(cat.ID, filesWithDesc)
		if err != nil {
			s.Logger.Errorf("error saving cat files: %v", err)
			return nil, err
		}
		cat.Files = &filesCat

	} else {
		s.Logger.Info("No files to save for this cat")
	}

	s.Logger.Infof("Service CreateCatShowCat OK")
	return cat, nil
}
