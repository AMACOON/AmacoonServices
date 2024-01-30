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
		Preload("Federation.Country").
		Preload("Cattery.Owner").
		Preload("Cattery.Country").
		Preload("Federation").
		Preload("Titles.Titles").
		Preload("Files").
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
	
	var catInfos []CatInfo

	err := r.DB.Table("cats").
		Select("cats.id, cats.name, breeds.breed_name as breed, colors.name as color, colors.ems_code, cats.neutered"). // Adicionando cats.neutered ao SELECT
		Joins("LEFT JOIN breeds ON breeds.id = cats.breed_id").
		Joins("LEFT JOIN colors ON colors.id = cats.color_id").
		Where("cats.owner_id = ?", ownerId).
		Scan(&catInfos).Error

	if err != nil {
		return nil, err
	}

	return catInfos, nil
}


func (r *CatRepository) UpdateNeuteredStatus(catID string, neutered bool) error {
	r.Logger.Infof("Repository UpdateNeuteredStatus")

	// Locate the record for the cat with the given ID
	cat := Cat{}
	if err := r.DB.First(&cat, catID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No cat found with id: %s", catID)
			return err
		}
		return err
	}

	// Update the Neutered status
	if err := r.DB.Model(&cat).Update("neutered", neutered).Error; err != nil {
		r.Logger.Errorf("Update Cat Neutered status failed: %v", err)
		return err
	}

	r.Logger.Infof("Repository UpdateNeuteredStatus OK")

	return nil
}

func (r *CatRepository) ValidateCat(name string, microchip string, registration string, registrationType string) (*Cat, error) {
	r.Logger.Infof("Repository FindCat")

	// Locate the record for the cat with the given parameters
	var cat Cat
	// Validar Regra com a Kleyne
	if err := r.DB.Where("name = ? AND microchip = ? AND registration = ? AND registrationType = ?", name, microchip, registration, registrationType).First(&cat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Infof("No cat found with name: %s, microchip: %s, registration: %s, registrationType: %s", name, microchip, registration, registrationType)
			return nil, nil
		}
		return nil, err
	}

	r.Logger.Infof("Repository FindCat OK")

	return &cat, nil
}

func (r *CatRepository) GetAllCats(filter string) ([]CatInfoAdm, error) {
	r.Logger.Infof("Repository GetAllCats")

	var results []CatInfoAdm

	db := r.DB.Table("cats").
		Select("cats.name, breeds.name as breed_name, owners.name as owner_name").
		Joins("left join breeds on cats.breed_id = breeds.id").
		Joins("left join owners on cats.owner_id = owners.id")

		switch filter {
		case "non_validated":
			db = db.Where("cats.validated = ?", false)
		case "blank_microchip":
			db = db.Where("cats.microchip = ?", "")
		case "blank_register":
			db = db.Where("cats.register = ?", "")
		case "blank_cattery":
			db = db.Where("cats.cattery_id is NULL")
		default:
			// do nothing, return all entries
		}
	

	result := db.Scan(&results)
	if result.Error != nil {
		r.Logger.Errorf("Error getting cats: %v", result.Error)
		return nil, result.Error
	}

	r.Logger.Infof("Repository GetAllCats OK")
	return results, nil
}

func (r *CatRepository) UpdateCat(id string, updatedCat *Cat) error {
	r.Logger.Infof("Repository UpdateCat")

	// Locate the record for the cat with the given ID
	cat := Cat{}
	if err := r.DB.First(&cat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No cat found with id: %s", id)

			return err
		}
		return err
	}

	// Start a new DB transaction
	tx := r.DB.Begin()

	// Update the Cat's fields
	if err := tx.Model(&cat).Updates(updatedCat).Error; err != nil {
		tx.Rollback()
		r.Logger.Errorf("Update Cat failed: %v", err)
		return err
	}

	// If the Titles field of updatedCat is not nil, update the titles
	if updatedCat.Titles != nil {
		for _, title := range *updatedCat.Titles {
			// If title ID is 0, it's a new title, create it
			if title.ID == 0 {
				if err := tx.Model(&TitlesCat{}).Create(&title).Error; err != nil {
					tx.Rollback()
					r.Logger.Errorf("Create title failed: %v", err)
					return err
				}
			} else { // Otherwise, update existing title
				if err := tx.Model(&TitlesCat{}).Where("id = ?", title.ID).Updates(&title).Error; err != nil {
					tx.Rollback()
					r.Logger.Errorf("Update title failed: %v", err)
					return err
				}
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.Errorf("Transaction commit failed: %v", err)
		return err
	}

	r.Logger.Infof("Repository UpdateCat OK")

	return nil
}








