package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

type BreedRepository struct {
	DB *gorm.DB
}

func (r *BreedRepository) GetAllBreeds() ([]models.Breed, error) {
	var breeds []models.Breed

	query := r.DB.Unscoped().Find(&breeds)
	if err := query.Error; err != nil {
		return nil, err
	}
	return breeds, nil
}

func (r *BreedRepository) GetCompatibleBreeds(BreedID string) ([]string, error) {
    var compatibleRaces []string

    query := r.DB.Unscoped().Table("racas_compat").
        Select("id_racas2").
        Where("id_racas1 = ?", BreedID).
        Find(&compatibleRaces)
    if query.Error != nil {
        return nil, query.Error
    }

    

    return compatibleRaces, nil
}

