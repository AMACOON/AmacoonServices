package owner

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"errors"
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
	if err := r.DB.Delete(&Owner{}, id).Error; err != nil {
		r.Logger.WithError(err).Errorf("error deleting owner by id: %v", id)
		return err
	}
	r.Logger.Infof("Repository DeleteOwnerByID OK")
	return nil
}

func (r *OwnerRepository) Login(loginRequest LoginRequest) (*Owner, error) {
	r.Logger.Info("Repository Login")
	
	user := &Owner{}
	if err := r.DB.Where("email = ?", loginRequest.Email).First(user).Error; err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect %v", loginRequest.Email)
		return nil, err
	}

	// Comparar a senha fornecida com a senha armazenada
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		r.Logger.WithError(err).Errorf("User not found or password incorrect")
		return nil, err
	}

	// A senha está correta, retornar o usuário
	r.Logger.Info("Repository Login OK")
	return user, nil
}

