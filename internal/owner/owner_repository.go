package owner

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
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
	if err := r.DB.First(&owner, id).Error; err != nil {
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

func (r *OwnerRepository) UpdateOwner(id uint, owner *Owner) (*Owner, error) {
    r.Logger.Infof("Repository UpdateOwner")
    var existingOwner Owner
    if err := r.DB.First(&existingOwner, id).Error; err != nil {
        r.Logger.WithError(err).Errorf("error finding owner with id %v", id)
        return nil, err
    }
    existingOwner.Name = owner.Name
    existingOwner.CPF = owner.CPF
    existingOwner.Address = owner.Address
    existingOwner.City = owner.City
    existingOwner.State = owner.State
    existingOwner.ZipCode = owner.ZipCode
    existingOwner.CountryID = owner.CountryID
    existingOwner.Phone = owner.Phone
    existingOwner.Valid = owner.Valid
    existingOwner.ValidId = owner.ValidId
    existingOwner.Observation = owner.Observation
    if err := r.DB.Save(&existingOwner).Error; err != nil {
        r.Logger.WithError(err).Errorf("error updating owner with id %v", id)
        return nil, err
    }
    r.Logger.Infof("Repository UpdateOwner OK")
    return &existingOwner, nil
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

