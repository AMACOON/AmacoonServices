package catshowcat

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatShowCatFileService struct {
	FileService   *utils.FilesService
	Logger        *logrus.Logger
	FilesCatRepo  *FilesCatShowCatRepository
}

func NewCatShowCatFileService(fileService *utils.FilesService, filesCatRepo *FilesCatShowCatRepository, logger *logrus.Logger) *CatShowCatFileService {
	return &CatShowCatFileService{
		FileService:  fileService,
		FilesCatRepo: filesCatRepo,
		Logger:       logger,
	}
}

func (s *CatShowCatFileService) SaveCatShowCatFiles(catShowCatID uint, filesWithDesc []utils.FileWithDescription) ([]FilesCatShowCat, error) {
	s.Logger.Infof("Service SaveCatShowCatFiles")
	
	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(catShowCatID), 10), "cats", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesCat and save them
	var filesCatCreated []FilesCatShowCat
	for _, file := range files {
		fileCat := FilesCatShowCat{
			CatShowCatID:    catShowCatID,
			FileData: file,
		}
		created, err := s.FilesCatRepo.CreateFilesCatShowCat([]FilesCatShowCat{fileCat})
		if err != nil {
			s.Logger.Errorf("Failed to create file cat: %v", err)
			return nil, err
		}
		filesCatCreated = append(filesCatCreated, created...)
	}

	s.Logger.Infof("Service SaveCatShowCatFiles OK")
	return filesCatCreated, nil
}
