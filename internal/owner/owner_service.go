package owner

import (
	"github.com/sirupsen/logrus"
    "strconv"
	"golang.org/x/crypto/bcrypt"
)

type OwnerService struct {
	OwnerRepo *OwnerRepository
	Logger    *logrus.Logger
}

func NewOwnerService(ownerRepo *OwnerRepository, logger *logrus.Logger) *OwnerService {
	return &OwnerService{
		OwnerRepo: ownerRepo,
		Logger:    logger,
	}
}

func (s *OwnerService) GetOwnerByID(idStr string) (*Owner, error) {
	s.Logger.Info("Service GetOwnerByID")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		s.Logger.WithError(err).Errorf("Invalid owner ID: %s", idStr)
		return nil, err
	}
	owner, err := s.OwnerRepo.GetOwnerByID(uint(id))
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get owner from repository")
		return nil, err
	}
	s.Logger.Info("Service GetOwnerByID OK")
	return owner, nil
}
func (s *OwnerService) GetAllOwners() ([]Owner, error) {
	s.Logger.Infof("Service GetAllOwners")
	owners, err := s.OwnerRepo.GetAllOwners()
	if err != nil {
		s.Logger.WithError(err).Error("failed to get all owners")
		return nil, err
	}
	s.Logger.Infof("Service GetAllOwners OK")
	return owners, nil
}

func (s *OwnerService) GetOwnerByCPF(cpf string) (*Owner, error) {
	s.Logger.Infof("Service GetOwnerByCPF")
	owner, err := s.OwnerRepo.GetOwnerByCPF(cpf)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get owner by CPF")
		return nil, err
	}
	s.Logger.Infof("Service GetOwnerByCPF OK")
	return owner, nil
}



func (s *OwnerService) CreateOwner(owner *Owner) (*Owner, error) {
	s.Logger.Infof("Service CreateOwner")

	// Gerar o hash da senha antes de salvar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(owner.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.WithError(err).Error("failed to generate password hash")
		return nil, err
	}

	owner.PasswordHash = string(hashedPassword) // Atualizar o campo PasswordHash com o hash gerado
	owner.Valid = false

	createdOwner, err := s.OwnerRepo.CreateOwner(owner)
	if err != nil {
		s.Logger.WithError(err).Error("failed to create owner")
		return nil, err
	}

	s.Logger.Infof("Service CreateOwner OK")
	return createdOwner, nil
}


func (s *OwnerService) UpdateOwner(idStr string, updatedOwner *Owner) error {
	s.Logger.Infof("Service UpdateOwner")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		s.Logger.WithError(err).Errorf("Invalid owner ID: %s", idStr)
		return err
	}
	err = s.OwnerRepo.UpdateOwner(uint(id), updatedOwner)
	if err != nil {
		s.Logger.WithError(err).Error("failed to update owner")
		return err
	}
	s.Logger.Infof("Service UpdateOwner OK")
	return nil
}


func (s *OwnerService) DeleteOwnerByID(idStr string) error {
	s.Logger.Infof("Service DeleteOwnerByID")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		s.Logger.WithError(err).Errorf("Invalid owner ID: %s", idStr)
		return err
	}
	err = s.OwnerRepo.DeleteOwnerByID(uint(id))
	if err != nil {
		s.Logger.WithError(err).Error("failed to delete owner")
		return err
	}
	s.Logger.Infof("Service DeleteOwnerByID OK")
	return nil
}

func (s *OwnerService) Login(loginRequest LoginRequest) (*Owner, error) {
	s.Logger.Info("Service Login")

	user, err := s.OwnerRepo.Login(loginRequest)
	if err != nil {
		return nil, err
	}
	
	s.Logger.Info("Service Login OK")
	return user, nil
}





