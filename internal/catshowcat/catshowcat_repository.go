package catshowcat

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatShowCatRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatShowCatRepository(db *gorm.DB, logger *logrus.Logger) *CatShowCatRepository {
	return &CatShowCatRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatShowCatRepository) CreateCatShowCat(cat *CatShowCat) (*CatShowCat, error) {
	r.Logger.Infof("Repository CreateCatShowCat")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the Cat record
	if err := tx.Create(cat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// If everything goes well, commit the transaction
	tx.Commit()

	r.Logger.Infof("Repository CreateCatShowCat OK")
	return cat, nil
}
