package catshowcat

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesCatShowCatRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesCatShowCatRepository(db *gorm.DB, logger *logrus.Logger) *FilesCatShowCatRepository {
	return &FilesCatShowCatRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesCatShowCatRepository) CreateFilesCatShowCat(filesCatShowCat []FilesCatShowCat) ([]FilesCatShowCat, error) {
	r.Logger.Infof("Repository CreateFilesCatShowCat")
	
	var filesCatCreated []FilesCatShowCat
	for _, fileCat := range filesCatShowCat {
		err := r.DB.Create(&fileCat).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file cat: %v", err)
			return nil, err
		}
		filesCatCreated = append(filesCatCreated, fileCat)
	}
	
	r.Logger.Infof("Repository CreateFilesCatShowCat OK")
	return filesCatCreated, nil
}
