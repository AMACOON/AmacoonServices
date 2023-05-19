package litter

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesLitterRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesLitterRepository(db *gorm.DB, logger *logrus.Logger) *FilesLitterRepository {
	return &FilesLitterRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesLitterRepository) CreateFilesLitter(filesLitter []FilesLitter) ([]FilesLitter, error) {
	r.Logger.Infof("Repository CreateFilesLitter")
	
	var filesLitterCreated []FilesLitter
	for _, fileLitter := range filesLitter {
		err := r.DB.Create(&fileLitter).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file litter: %v", err)
			return nil, err
		}
		filesLitterCreated = append(filesLitterCreated, fileLitter)
	}
	
	r.Logger.Infof("Repository CreateFilesCat OK")
	return filesLitterCreated, nil
}
