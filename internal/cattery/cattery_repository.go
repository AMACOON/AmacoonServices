package cattery

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatteryRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatteryRepository(db *gorm.DB, logger *logrus.Logger) *CatteryRepository {
	return &CatteryRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatteryRepository) GetCatteryByID(id string) (*Cattery, error) {
	r.Logger.Infof("Repository GetCatteryByID")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err
	}
	var cattery Cattery
	if err := r.DB.First(&cattery, uintID).Error; err != nil {
		r.Logger.WithError(err).Errorf("failed to get cattery with ID %s", id)
		return nil, err
	}
	r.Logger.Infof("Repository GetCatteryByID OK")
	return &cattery, nil
}

func (r *CatteryRepository) GetAllCatteries() ([]Cattery, error) {
	r.Logger.Infof("Repository GetAllCatteries")
	var catteries []Cattery
	if err := r.DB.Find(&catteries).Error; err != nil {
		r.Logger.WithError(err).Errorf("failed to get all catteries")
		return nil, err
	}
	r.Logger.Infof("Repository GetAllCatteries OK")
	return catteries, nil
}
