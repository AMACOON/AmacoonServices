package breed

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BreedRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewBreedRepository(db *gorm.DB, logger *logrus.Logger) *BreedRepository {
	return &BreedRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *BreedRepository) GetBreedByID(id uint) (*Breed, error) {
	r.Logger.Infof("Repository GetBreedByID")

	var breed Breed
	result := r.DB.First(&breed, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("breed not found")
		}
		return nil, result.Error
	}

	r.Logger.Infof("Repository GetBreedByID OK")
	return &breed, nil
}

func (r *BreedRepository) GetAllBreeds() ([]Breed, error) {
	r.Logger.Infof("Repository GetAllBreeds")

	var breeds []Breed

	result := r.DB.Find(&breeds)
	if result.Error != nil {
		r.Logger.WithError(result.Error).Error("failed to get all breeds")
		return nil, result.Error
	}

	r.Logger.Infof("Repository GetAllBreeds OK")
	return breeds, nil
}
