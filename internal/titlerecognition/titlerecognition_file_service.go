package titlerecognition

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type FilesTitleRecognitionService struct {
	FileService *utils.FilesService
	Logger      *logrus.Logger
	FilesTitleRecognitionRepo     *FilesTitleRecognitionRepository
}

func NewFilesTitleRecognitionService(fileService *utils.FilesService, filesTitleRecognitionRepo *FilesTitleRecognitionRepository, logger *logrus.Logger) *FilesTitleRecognitionService {
	return &FilesTitleRecognitionService{
		FileService:  fileService,
		FilesTitleRecognitionRepo: filesTitleRecognitionRepo,
		Logger:       logger,
	}
}

func (s *FilesTitleRecognitionService) SaveTitleRecognitionFiles(TitleRecognitionID uint, filesWithDesc []utils.FileWithDescription) ([]FilesTitleRecognition, error) {
	s.Logger.Infof("Service SaveTitleRecognitionFiles")

	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(TitleRecognitionID), 10), "titlerecognition", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error save TitlesRecognition files: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesLitter and save them
	var filesTitleRecognitionCreated []FilesTitleRecognition
	for _, file := range files {
		fileTitleRecognition := FilesTitleRecognition{
			TitleRecognitionID:    TitleRecognitionID,
			FileData: file,
		}
		created, err := s.FilesTitleRecognitionRepo.CreateFilesTitleRecognition([]FilesTitleRecognition{fileTitleRecognition})
		if err != nil {
			s.Logger.Errorf("Failed to create file TitlesRecognition Insert table: %v", err)
			return nil, err
		}
		filesTitleRecognitionCreated = append(filesTitleRecognitionCreated, created...)
	}

	s.Logger.Infof("Service SaveTitleRecognitionFiles OK")
	return filesTitleRecognitionCreated, nil
}
