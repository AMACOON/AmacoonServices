package cat

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatRepository(db *gorm.DB, logger *logrus.Logger) *CatRepository {
	return &CatRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatRepository) GetCatCompleteByID(id string) (*Cat, error) {

	r.Logger.Infof("Repository GetCatCompleteByID")
	var cat Cat

	result := r.DB.
		Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner.Country").
		Preload("Federation").
		Preload("Titles.Titles").
		Where("id = ?", id).First(&cat)

	if cat.FatherID != nil {
		var father Cat
		r.DB.Select("name").Where("id = ?", cat.FatherID).First(&father)
		cat.FatherName = father.Name
	}

	if cat.MotherID != nil {
		var mother Cat
		r.DB.Select("name").Where("id = ?", cat.MotherID).First(&mother)
		cat.MotherName = mother.Name
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			r.Logger.Infof("Repository GetCatCompleteByID not found")
			return nil, nil
		}
		return nil, result.Error
	}
	r.Logger.Infof("Repository GetCatCompleteByID OK")
	return &cat, nil
}

func (r *CatRepository) GetCatCompleteByAllByOwner(ownerID string) ([]*Cat, error) {
	r.Logger.Infof("Repository GetCatCompleteByAllByOwner")

	var catsComplete []*Cat
	query := r.DB.
		Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner.Country").
		Preload("Federation").
		Preload("Titles.Titles")
	err := query.Where("owner_id = ?", ownerID).Find(&catsComplete).Error

	if err != nil {
		r.Logger.WithError(err).Error("error getting cats")
		return nil, err
	}

	for _, cat := range catsComplete {
		if cat.FatherID != nil {
			var father Cat
			r.DB.Select("name").Where("id = ?", cat.FatherID).First(&father)
			cat.FatherName = father.Name
		}

		if cat.MotherID != nil {
			var mother Cat
			r.DB.Select("name").Where("id = ?", cat.MotherID).First(&mother)
			cat.MotherName = mother.Name
		}
	}

	r.Logger.Infof("Repository GetCatCompleteByAllByOwner OK")
	return catsComplete, nil
}
