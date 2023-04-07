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



