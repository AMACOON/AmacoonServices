package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProtocolRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewProtocolRepository(db *gorm.DB, logger *logrus.Logger) *ProtocolRepository {
	return &ProtocolRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *ProtocolRepository) ProtocolNumberExists(protocol string) (bool, error) {
	r.Logger.Infof("Repository ProtocolNumberExists")
	var count int64
	err := r.DB.Model(&Protocol{}).Where("protocol = ?", protocol).Count(&count).Error
	if err != nil {
		return false, err
	}
	r.Logger.Infof("Repository ProtocolNumberExists OK")
	return count > 0, nil
}

func (r *ProtocolRepository) SaveProtocolNumber(protocolNumber string) error {
	r.Logger.Infof("Repository SaveProtocolNumber")
	protocol := Protocol{
		Protocol: protocolNumber,
	}
	err := r.DB.Create(&protocol).Error
	if err != nil {
		return err
	}
	r.Logger.Infof("Repository SaveProtocolNumber OK")
	return nil
}
