package titlerecognition

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesTitleRecognitionRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesTitleRecognitionRepository(db *gorm.DB, logger *logrus.Logger) *FilesTitleRecognitionRepository {
	return &FilesTitleRecognitionRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesTitleRecognitionRepository) CreateFilesTitleRecognition(filesTitleRecognition []FilesTitleRecognition) ([]FilesTitleRecognition, error) {
	r.Logger.Infof("Repository CreateFilesTitleRecognition")
	
	var filesTitleRecognitionCreated []FilesTitleRecognition
	for _, fileTitleRecognition := range filesTitleRecognition {
		err := r.DB.Create(&fileTitleRecognition).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file TitleRecognition: %v", err)
			return nil, err
		}
		filesTitleRecognitionCreated = append(filesTitleRecognitionCreated, fileTitleRecognition)
	}
	
	r.Logger.Infof("Repository CreateFilesTitleRecognition OK")
	return filesTitleRecognitionCreated, nil
}
