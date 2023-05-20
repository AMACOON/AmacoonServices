package owner

import (
	"strconv"

	"fmt"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type OwnerService struct {
	OwnerRepo         *OwnerRepository
	OwnerEmailService *OwnerEmailService
	Logger            *logrus.Logger
}

func NewOwnerService(ownerRepo *OwnerRepository, ownerEmailService *OwnerEmailService, logger *logrus.Logger) *OwnerService {
	return &OwnerService{
		OwnerRepo:         ownerRepo,
		OwnerEmailService: ownerEmailService,
		Logger:            logger,
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

	// Verificar se já existe um proprietário com o mesmo nome, e-mail ou CPF
	exists, err := s.OwnerRepo.CheckOwnerExistence(owner.Name, owner.Email, owner.CPF)
	if err != nil {
		s.Logger.WithError(err).Error("failed to check owner existence")
		return nil, err
	}
	if exists {
		err := fmt.Errorf("owner with the same name, email, or CPF already exists")
		s.Logger.Errorf("owner with same name, email, or CPF already exists")
		return nil, err
	}

	// Gerar o hash da senha antes de salvar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(owner.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.WithError(err).Error("failed to generate password hash")
		return nil, err
	}

	// Gerar o random validação antes de salvar
	randomString := utils.GenerateRandomString()

	// Atualizar o campo PasswordHash com o hash gerado e o campo ValidId com o random gerado
	owner.PasswordHash = string(hashedPassword)
	owner.ValidId = randomString
	owner.Valid = false

	_, err = s.OwnerRepo.CreateOwner(owner)
	if err != nil {
		s.Logger.WithError(err).Error("failed to create owner")
		return nil, err
	}

	// Consultar o proprietário criado para obter Informações adicionais
	createdOwner, err := s.GetOwnerByID(strconv.Itoa(int(owner.ID)))
	if err != nil {
		s.Logger.WithError(err).Error("failed to get owner by ID")
		return nil, err
	}
	// Enviar email de validação
	err = s.OwnerEmailService.SendOwnerValidationEmail(createdOwner)
	if err != nil {
		s.Logger.WithError(err).Error("failed to send validation email")
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
	//Pegar dados do Owener para verificar se a senha mudou, se mudou gerar o hash
	owner, err := s.OwnerRepo.GetOwnerByID(uint(id))
	if err != nil {
		s.Logger.WithError(err).Error("failed to get owner by ID")
		return err
	}
	if owner.PasswordHash != updatedOwner.PasswordHash {
		// Gerar o hash da senha antes de salvar
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedOwner.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			s.Logger.WithError(err).Error("failed to generate password hash")
			return err
		}
		updatedOwner.PasswordHash = string(hashedPassword)
	}
	// Atualizar os campos do proprietário na tabela "owners"
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

func (s *OwnerService) UpdateValidOwner(id string, validID string) error {
	s.Logger.Infof("Service UpdateValidOwner")

	ownerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		s.Logger.WithError(err).Error("failed to parse owner ID")
		return err
	}

	err = s.OwnerRepo.UpdateValidOwner(uint(ownerID), validID)
	if err != nil {
		s.Logger.WithError(err).Error("failed to update owner")
		return err
	}
	// Consultar o proprietário criado para obter Informações adicionais
	updateOwner, err := s.GetOwnerByID(strconv.Itoa(int(ownerID)))
	if err != nil {
		s.Logger.WithError(err).Error("failed to get owner by ID")
		return err
	}
	// Enviar email de confirmação
	err = s.OwnerEmailService.SendOwnerValidationConfirmationEmail(updateOwner)
	if err != nil {
		s.Logger.WithError(err).Error("failed to send validation email")
		return err
	}

	s.Logger.Infof("Service UpdateValidOwner OK")
	return nil
}
