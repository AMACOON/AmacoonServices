package cattery

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesCatteryRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesCatteryRepository(db *gorm.DB, logger *logrus.Logger) *FilesCatteryRepository {
	return &FilesCatteryRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesCatteryRepository) CreateFilesCattery(filesCattery []FilesCattery) ([]FilesCattery, error) {
	r.Logger.Infof("Repository CreateFilesCattery")
	
	var filesCatteryCreated []FilesCattery
	for _, fileCattery := range filesCattery {
		err := r.DB.Create(&fileCattery).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file Cattery: %v", err)
			return nil, err
		}
		filesCatteryCreated = append(filesCatteryCreated, fileCattery)
	}
	
	r.Logger.Infof("Repository CreateFilesCattery OK")
	return filesCatteryCreated, nil
}
