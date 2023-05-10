package title

import (
	"fmt"

	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

type TitleRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewTitleRepository(db *gorm.DB, logger *logrus.Logger) *TitleRepository {
	return &TitleRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *TitleRepository) GetAllTitles() ([]Title, error) {
	r.Logger.Infof("Repository GetAllTitles")
	var titles []Title
	result := r.DB.Find(&titles)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching titles: %v", result.Error)
	}
	r.Logger.Infof("Repository GetAllTitles OK")
	return titles, nil
}
