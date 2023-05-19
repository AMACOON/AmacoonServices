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


func (r *CatRepository) CreateCat(cat *Cat) (*Cat, error) {
	r.Logger.Infof("Repository CreateCat")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the Cat record
	if err := tx.Create(&cat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// If everything goes well, commit the transaction
	tx.Commit()

	r.Logger.Infof("Repository CreateCat OK")
	return cat, nil
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
			r.Logger.Errorf("Repository GetCatCompleteByID not found")
			return nil, nil
		}
		return nil, result.Error
	}
	r.Logger.Infof("Repository GetCatCompleteByID OK")
	return &cat, nil
}

func (r *CatRepository) GetCatsByOwner(ownerId string) ([]CatInfo, error) {
	r.Logger.Infof("Repository GetCatsByOwner")
	
	var cats []Cat
	var catInfos []CatInfo

	// Make sure to preload the needed fields to avoid lazy loading
	if err := r.DB.Where("owner_id = ?", ownerId).Preload("Breed").Preload("Color").Find(&cats).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No cats found for owner id: %s", ownerId)
			return nil, err
		}
		return nil, err
	}

	for _, cat := range cats {
		catInfo := CatInfo{
			Name:  cat.Name,
			Breed: cat.Breed.BreedName, 
			Color: cat.Color.Name,
		}

		catInfos = append(catInfos, catInfo)
	}

	return catInfos, nil
}

func (r *CatRepository) UpdateNeuteredStatus(catID string, neutered bool) error {
	r.Logger.Infof("Repository UpdateNeuteredStatus")

	// Locate the record for the cat with the given ID
	cat := Cat{}
	if err := r.DB.First(&cat, catID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No cat found with id: %d", catID)
			return err
		}
		return err
	}

	// Update the Neutered status
	if err := r.DB.Model(&cat).Update("Neutered", neutered).Error; err != nil {
		r.Logger.Errorf("Update Cat Neutered status failed: %v", err)
		return err
	}

	r.Logger.Infof("Repository UpdateNeuteredStatus OK")

	return nil
}


