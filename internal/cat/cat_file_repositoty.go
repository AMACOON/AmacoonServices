package cat

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesCatRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesCatRepository(db *gorm.DB, logger *logrus.Logger) *FilesCatRepository {
	return &FilesCatRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesCatRepository) CreateFilesCat(filesCat []FilesCat) ([]FilesCat, error) {
	r.Logger.Infof("Repository CreateFilesCat")
	
	var filesCatCreated []FilesCat
	for _, fileCat := range filesCat {
		err := r.DB.Create(&fileCat).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file cat: %v", err)
			return nil, err
		}
		filesCatCreated = append(filesCatCreated, fileCat)
	}
	
	r.Logger.Infof("Repository CreateFilesCat OK")
	return filesCatCreated, nil
}
