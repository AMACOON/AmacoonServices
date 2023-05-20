package cat

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatFileService struct {
	FileService   *utils.FilesService
	Logger        *logrus.Logger
	FilesCatRepo  *FilesCatRepository
}

func NewCatFileService(fileService *utils.FilesService, filesCatRepo *FilesCatRepository, logger *logrus.Logger) *CatFileService {
	return &CatFileService{
		FileService:  fileService,
		FilesCatRepo: filesCatRepo,
		Logger:       logger,
	}
}

func (s *CatFileService) SaveCatFiles(catID uint, filesWithDesc []utils.FileWithDescription) ([]FilesCat, error) {
	s.Logger.Infof("Service SaveCatFiles")
	
	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(catID), 10), "cats", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesCat and save them
	var filesCatCreated []FilesCat
	for _, file := range files {
		fileCat := FilesCat{
			CatID:    catID,
			FileData: file,
		}
		created, err := s.FilesCatRepo.CreateFilesCat([]FilesCat{fileCat})
		if err != nil {
			s.Logger.Errorf("Failed to create file cat: %v", err)
			return nil, err
		}
		filesCatCreated = append(filesCatCreated, created...)
	}

	s.Logger.Infof("Service SaveCatFiles OK")
	return filesCatCreated, nil
}
