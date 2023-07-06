package cattery

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatteryFileService struct {
	FileService   *utils.FilesService
	Logger        *logrus.Logger
	FilesCatteryRepo  *FilesCatteryRepository
}

func NewCatteryFileService(fileService *utils.FilesService, filesCatteryRepo *FilesCatteryRepository, logger *logrus.Logger) *CatteryFileService {
	return &CatteryFileService{
		FileService:  fileService,
		FilesCatteryRepo: filesCatteryRepo,
		Logger:       logger,
	}
}

func (s *CatteryFileService) SaveCatteryFiles(CatteryID uint, filesWithDesc []utils.FileWithDescription) ([]FilesCattery, error) {
	s.Logger.Infof("Service SaveCatteryFiles")
	
	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(CatteryID), 10), "catteries", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error creating Cattery from repository: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesCattery and save them
	var filesCatteryCreated []FilesCattery
	for _, file := range files {
		fileCattery := FilesCattery{
			CatteryID:    CatteryID,
			FileData: file,
		}
		created, err := s.FilesCatteryRepo.CreateFilesCattery([]FilesCattery{fileCattery})
		if err != nil {
			s.Logger.Errorf("Failed to create file Cattery: %v", err)
			return nil, err
		}
		filesCatteryCreated = append(filesCatteryCreated, created...)
	}

	s.Logger.Infof("Service SaveCatteryFiles OK")
	return filesCatteryCreated, nil
}
