package transfer

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	
)

type FilesTransferRepository struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewFilesTransferRepository(db *gorm.DB, logger *logrus.Logger) *FilesTransferRepository {
	return &FilesTransferRepository{
		Logger: logger,
		DB:     db,
	}
}

func (r *FilesTransferRepository) CreateFilesTransfer(filesTransfer []FilesTransfer) ([]FilesTransfer, error) {
	r.Logger.Infof("Repository CreateFilesTransfer")
	
	var filesTransferCreated []FilesTransfer
	for _, fileTransfer := range filesTransfer {
		err := r.DB.Create(&fileTransfer).Error
		if err != nil {
			r.Logger.Errorf("Failed to create file Transfer: %v", err)
			return nil, err
		}
		filesTransferCreated = append(filesTransferCreated, fileTransfer)
	}
	
	r.Logger.Infof("Repository CreateFilesCat OK")
	return filesTransferCreated, nil
}
