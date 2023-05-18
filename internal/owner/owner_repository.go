package owner

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OwnerRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewOwnerRepository(db *gorm.DB, logger *logrus.Logger) *OwnerRepository {
	return &OwnerRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *OwnerRepository) GetOwnerByID(id uint) (*Owner, error) {
	r.Logger.Infof("Repository GetOwnerByID")
	var owner Owner
	if err := r.DB.
		Preload("Clubs").
		Preload("Clubs.Club").
		Preload("Country").
		First(&owner, id).Error; err != nil {
		r.Logger.WithError(err).Errorf("error getting owner by id: %v", id)
		return nil, err
	}
	r.Logger.Infof("Repository GetOwnerByID OK")
	return &owner, nil
}

func (r *OwnerRepository) GetAllOwners() ([]Owner, error) {
	r.Logger.Infof("Repository GetAllOwners")
	var owners []Owner
	if err := r.DB.Find(&owners).Error; err != nil {
		r.Logger.WithError(err).Errorf("error getting all owners")
		return nil, err
	}
	r.Logger.Infof("Repository GetAllOwners OK")
	return owners, nil
}

func (r *OwnerRepository) GetOwnerByCPF(cpf string) (*Owner, error) {
	r.Logger.Infof("Repository GetOwnerByCPF")
	var owner Owner
	if err := r.DB.Where("cpf = ?", cpf).First(&owner).Error; err != nil {
		r.Logger.WithError(err).Errorf("error getting owner by CPF: %v", cpf)
		return nil, err
	}
	r.Logger.Infof("Repository GetOwnerByCPF OK")
	return &owner, nil
}

func (r *OwnerRepository) CreateOwner(owner *Owner) (*Owner, error) {
	r.Logger.Infof("Repository CreateOwner")
	if err := r.DB.Create(owner).Error; err != nil {
		r.Logger.WithError(err).Errorf("error creating owner: %v", owner)
		return nil, err
	}
	r.Logger.Infof("Repository CreateOwner OK")
	return owner, nil
}

func (r *OwnerRepository) UpdateOwner(id uint, owner *Owner) error {
	r.Logger.Infof("Repository UpdateOwner")

	// Verificar se o registro existe
	var existingOwner Owner
	result := r.DB.First(&existingOwner, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.Logger.Errorf("owner with id %v not found", id)
			return result.Error
		}
		r.Logger.Errorf("error finding owner with id %v: %v", id, result.Error)
		return result.Error
	}

	// Atualizar os campos do proprietário na tabela "owners"
	if err := r.DB.Model(&existingOwner).Updates(owner).Error; err != nil {
		r.Logger.Errorf("error updating owner: %v", err)
		return err
	}

	// Atualizar os campos dos clubes relacionados na tabela "owners_clubs"
	for _, club := range owner.Clubs {
		if err := r.DB.Model(&OwnerClub{}).Where("id = ?", club.ID).Updates(club).Error; err != nil {
			r.Logger.Errorf("error updating owner club record with id %v: %v", club.ID, err)
			return err
		}
	}

	r.Logger.Infof("Repository UpdateOwner OK")
	return nil
}

func (r *OwnerRepository) DeleteOwnerByID(id uint) error {
	r.Logger.Infof("Repository DeleteOwnerByID")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var owner Owner
	if err := tx.First(&owner, id).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error finding owner with id %v", id)
		return err
	}

	// Delete the clubs associated with this owner from "owners_clubs"
	if err := tx.Where("owner_id = ?", id).Delete(&OwnerClub{}).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting owner club records with owner id %v", id)
		return err
	}

	// Delete the owner from "owners"
	if err := tx.Delete(&owner).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting owner with id %v", id)
		return err
	}

	// If everything goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.WithError(err).Errorf("error committing transaction")
		return err
	}

	r.Logger.Infof("Repository DeleteOwnerByID OK")
	return nil
}

func (r *OwnerRepository) CheckOwnerExistence(name, email, cpf string) (bool, error) {
	r.Logger.Infof("Repository CheckOwnerExistence")

	var count int64
	result := r.DB.Model(&Owner{}).Where("name = ? OR email = ? OR cpf = ?", name, email, cpf).Count(&count)
	if result.Error != nil {
		r.Logger.WithError(result.Error).Error("error checking owner existence")
		return false, result.Error
	}

	r.Logger.Infof("Repository CheckOwnerExistence OK")
	return count > 0, nil
}

func (r *OwnerRepository) UpdateValidOwner(id uint, validID string) error {
	r.Logger.Infof("Repository UpdateValidOwner")

	var existingOwner Owner
	result := r.DB.First(&existingOwner, id)
	if result.Error != nil {
		r.Logger.WithError(result.Error).Errorf("error finding owner with id %v", id)
		return result.Error
	}

	if existingOwner.ValidId != validID {
		r.Logger.Errorf("validation ID mismatch for owner with id %v", id)
		return result.Error
	}

	result = r.DB.Model(&Owner{}).Where("id = ?", id).Updates(map[string]interface{}{
		"valid": true,
	})
	if result.Error != nil {
		r.Logger.WithError(result.Error).Errorf("error updating owner with id %v", id)
		return result.Error
	}

	r.Logger.Infof("Repository UpdateValidOwner OK")
	return nil
}
