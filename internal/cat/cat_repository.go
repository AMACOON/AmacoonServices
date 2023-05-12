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

// func (r *CatRepository) GetAllByOwnerAndGender(ownerID, gender string) ([]*Cat, error) {
// 	r.Logger.Infof("Repository GetAllByOwnerAndGender")

// 	var catsComplete []*Cat
// 	query := r.DB.Preload("Color").Preload("Breed").Preload("Mother").Preload("Cattery").Preload("Country").Preload("Federation").Preload("Owner").Preload("Father").Preload("Titles")
// 	err := query.Where("owner_id = ? AND gender = ?", ownerID, gender).Find(&catsComplete).Error

// 	if err != nil {
// 		r.Logger.WithError(err).Error("error getting cats")
// 		return nil, err
// 	}

// 	r.Logger.Infof("Repository GetAllByOwnerAndGender OK")
// 	return catsComplete, nil
// }

// func (r *CatRepository) GetAllByOwner(ownerID string) ([]*Cat, error) {
// 	r.Logger.Infof("Repository GetAllByOwner")

// 	var catsComplete []*Cat
// 	query := r.DB.Preload("Breed").
// 		Preload("Color").
// 		Preload("Cattery").
// 		Preload("Country").
// 		Preload("Owner").
// 		Preload("Federation")
// 	err := query.Where("owner_id = ?", ownerID).Find(&catsComplete).Error

// 	if err != nil {
// 		r.Logger.WithError(err).Error("error getting cats")
// 		return nil, err
// 	}

// 	r.Logger.Infof("Repository GetAllByOwner OK")
// 	return catsComplete, nil
// }

// func (r *CatRepository) GetCatCompleteByRegistration(registration string) (*Cat, error) {
// 	r.Logger.Infof("Repository GetCatCompleteByRegistration")

// 	var catComplete Cat
// 	query := r.DB.Preload("Color").Preload("Breed").Preload("Mother").Preload("Cattery").Preload("Country").Preload("Federation").Preload("Owner").Preload("Father").Preload("Titles")
// 	err := query.Where("registration = ?", registration).First(&catComplete).Error

// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			r.Logger.WithField("Registration", registration).Warn("Cat not found")
// 		} else {
// 			r.Logger.WithError(err).Error("error getting cat")
// 		}
// 		return nil, err
// 	}

// 	r.Logger.Infof("Repository GetCatCompleteByRegistration OK")
// 	return &catComplete, nil
// }


