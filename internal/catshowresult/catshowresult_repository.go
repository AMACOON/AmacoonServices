package catshowresult

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatShowResultRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatShowResultRepository(db *gorm.DB, logger *logrus.Logger) *CatShowResultRepository {
	return &CatShowResultRepository{
		DB:     db,
		Logger: logger,
	}
}

func (repo *CatShowResultRepository) Create(catShowResult *CatShowResult) (*CatShowResult, error) {
	repo.Logger.Info("Creating new CatShowResult")
	if err := repo.DB.Create(catShowResult).Error; err != nil {
		repo.Logger.Errorf("Failed to create CatShowResult: %v", err)
		return nil, err
	}
	return catShowResult, nil
}


func (repo *CatShowResultRepository) GetById(id uint) (*CatShowResult, error) {
	repo.Logger.Infof("Fetching CatShowResult with ID: %d", id)
	var catShowResult CatShowResult
	if err := repo.DB.Preload("CatShowResultMatrix").First(&catShowResult, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			repo.Logger.Info("CatShowResult not found")
			return nil, nil
		}
		repo.Logger.Errorf("Failed to fetch CatShowResult: %v", err)
		return nil, err
	}
	return &catShowResult, nil
}

func (repo *CatShowResultRepository) GetByRegistrationID(registrationID uint) (*CatShowResult, error) {
    repo.Logger.Infof("Fetching CatShowResult with RegistrationID: %d", registrationID)
    var catShowResult CatShowResult
    if err := repo.DB.Where("registration_id = ?", registrationID).Preload("CatShowResultMatrix").First(&catShowResult).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            repo.Logger.Infof("CatShowResult not found for RegistrationID: %d", registrationID)
            return nil, nil
        }
        repo.Logger.Errorf("Failed to fetch CatShowResult by RegistrationID: %v", err)
        return nil, err
    }
    return &catShowResult, nil
}


func (repo *CatShowResultRepository) Update(id uint, catShowResult *CatShowResult) error {
	repo.Logger.Infof("Updating CatShowResult with ID: %d", id)
	if err := repo.DB.Model(&CatShowResult{}).Where("id = ?", id).Updates(catShowResult).Error; err != nil {
		repo.Logger.Errorf("Failed to update CatShowResult: %v", err)
		return err
	}
	return nil
}

func (repo *CatShowResultRepository) Delete(id uint) error {
	repo.Logger.Infof("Deleting CatShowResult with ID: %d", id)
	if err := repo.DB.Delete(&CatShowResult{}, id).Error; err != nil {
		repo.Logger.Errorf("Failed to delete CatShowResult: %v", err)
		return err
	}
	return nil
}
