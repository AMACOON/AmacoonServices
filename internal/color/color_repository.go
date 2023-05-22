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

// UpdateColor updates the color with the given ID with the new values
func (r *ColorRepository) UpdateColor(id string, updatedColor *Color) error {
	r.Logger.Infof("Repository UpdateColor")

	// Locate the record for the color with the given ID
	color := Color{}
	if err := r.DB.First(&color, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No color found with id: %d", id)
			return err
		}
		return err
	}

	// Update the color's fields
	if err := r.DB.Model(&color).Updates(updatedColor).Error; err != nil {
		r.Logger.Errorf("Update Color failed: %v", err)
		return err
	}

	r.Logger.Infof("Repository UpdateColor OK")

	return nil
}

// GetColorById retrieves a color by its ID
func (r *ColorRepository) GetColorById(id string) (*Color, error) {
	r.Logger.Infof("Repository GetColorById")

	color := &Color{}
	if err := r.DB.First(color, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No color found with id: %d", id)
			return nil, err
		}
		return nil, err
	}

	r.Logger.Infof("Repository GetColorById OK")

	return color, nil
}

