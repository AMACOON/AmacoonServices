package litter

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type FilesLitterService struct {
	FileService *utils.FilesService
	Logger      *logrus.Logger
	FilesLitterRepo     *FilesLitterRepository
}

func NewCatFileService(fileService *utils.FilesService, filesLitterRepo *FilesLitterRepository, logger *logrus.Logger) *FilesLitterService {
	return &FilesLitterService{
		FileService:  fileService,
		FilesLitterRepo: filesLitterRepo,
		Logger:       logger,
	}
}

func (s *FilesLitterService) SaveLitterFiles(litterID uint, filesWithDesc []utils.FileWithDescription) ([]FilesLitter, error) {
	s.Logger.Infof("Service SaveLitterFiles")

	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(litterID), 10), "litters", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesLitter and save them
	var filesLitterCreated []FilesLitter
	for _, file := range files {
		fileLitter := FilesLitter{
			LitterID:    litterID,
			FileData: file,
		}
		created, err := s.FilesLitterRepo.CreateFilesLitter([]FilesLitter{fileLitter})
		if err != nil {
			s.Logger.Errorf("Failed to create file litter: %v", err)
			return nil, err
		}
		filesLitterCreated = append(filesLitterCreated, created...)
	}

	s.Logger.Infof("Service SaveLitterFiles OK")
	return filesLitterCreated, nil
}
