package utils

import (
	

	"gorm.io/gorm"
)

type FilesRepository struct {
	DB *gorm.DB
}

func NewFilesRepository(db *gorm.DB) *FilesRepository {
	return &FilesRepository{
		DB: db,
	}
}

func (r *FilesRepository) GetFilesByServiceID(serviceID uint) ([]*FilesDB, error) {
	var files []*FilesDB
	if err := r.DB.Where("service_id = ?", serviceID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

/* func (r *FilesRepository) SaveServiceDocuments(files []models.Files) error {
	for _, file := range files {
		fileDB := models.FilesDB{
			Name:           file.Name,
			Type:           file.Type,
			Base64:         file.Base64,
			ProtocolNumber: file.ProtocolNumber,
			ServiceID:      file.ServiceID,
		}
		if err := r.DB.Create(&fileDB).Error; err != nil {
			return fmt.Errorf("failed to save service file: %w", err)
		}
	}
	return nil
} */

