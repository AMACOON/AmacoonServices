package color

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ColorRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewColorRepository(db *gorm.DB, logger *logrus.Logger) *ColorRepository {
	return &ColorRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *ColorRepository) GetAllColorsByBreed(breedCode string) ([]Color, error) {
	r.Logger.Infof("Repository GetAllColorsByBreed")
	var colors []Color
	result := r.DB.Where("breed_code = ?", breedCode).Find(&colors)
	if result.Error != nil {
		r.Logger.WithError(result.Error).Errorf("Failed to get colors for breed %s", breedCode)
		return nil, result.Error
	}
	r.Logger.Infof("Repository GetAllColorsByBreed OK")
	return colors, nil
}
