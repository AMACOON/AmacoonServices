package transfer

import (
	"errors"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type TransferService struct {
	TransferRepo      *TransferRepository
	ProtocolService   *utils.ProtocolService
	Logger            *logrus.Logger
}

func NewTransferService(transferRepo *TransferRepository, logger *logrus.Logger, protocolService *utils.ProtocolService) *TransferService {
	return &TransferService{
		TransferRepo:      transferRepo,
		ProtocolService:   protocolService,
		Logger:            logger,
	}
}

func (s *TransferService) CreateTransfer(req TransferRequest) (Transfer, error) {
	s.Logger.Infof("Service CreateTransfer")

	reqEntity, err := ConvertTransferRequestToTransfer(req)
	if err != nil {
		s.Logger.Errorf("error converting transfer request to transfer: %v", err)
		return Transfer{}, err
	}
	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("P")
	reqEntity.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generating protocol for transfer: %v", err)
		return Transfer{}, err
	}
	reqEntity.Status = "submitted"
	transfer, err := s.TransferRepo.CreateTransfer(reqEntity)
	if err != nil {
		s.Logger.Errorf("error fetching transfer from repository: %v", err)
		return Transfer{}, err
	}

	s.Logger.Infof("Service CreateTransfer OK")
	return transfer, nil
}

func (s *TransferService) GetTransferByID(id string) (*Transfer, error) {
	s.Logger.Infof("Service GetTransferByID")

	transfer, err := s.TransferRepo.GetTransferByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching transfer from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetTransferByID OK")
	return &transfer, nil
}

func (s *TransferService) UpdateTransferStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateTransferStatus")

	// check if the transfer exists
	if _, err := s.TransferRepo.GetTransferByID(id); err != nil {
		return errors.New("transfer not found")
	}

	if err := s.TransferRepo.UpdateTransferStatus(id, status); err != nil {
		return err
	}

	s.Logger.Infof("Service UpdateTransferStatus OK")
	return nil
}

func (s *TransferService) AddTransferFiles(id string, files []utils.Files) error {
	s.Logger.Infof("Service AddTransferFiles")

	// check if the transfer exists
	if _, err := s.TransferRepo.GetTransferByID(id); err != nil {
		return errors.New("transfer not found")
	}

	if err := s.TransferRepo.AddTransferFiles(id, files); err != nil {
		return err
	}

	s.Logger.Infof("Service AddTransferFiles OK")
	return nil
}

func (s *TransferService) GetTransferFilesByID(id string) ([]utils.Files, error) {
	s.Logger.Infof("Service GetTransferFilesByID")

	// check if the transfer exists
	if _, err := s.TransferRepo.GetTransferByID(id); err != nil {
		return nil, errors.New("transfer not found")
	}

	files, err := s.TransferRepo.GetTransferFilesByID(id)
	if err != nil {
		return nil, err
	}

	s.Logger.Infof("Service GetTransferFilesByID OK")
	return files, nil
}

func (s *TransferService) GetAllTransfersByRequesterID(requesterID string) ([]Transfer, error) {
	s.Logger.Infof("Service GetAllTransfersByRequesterID")
	transfers, err := s.TransferRepo.GetAllTransfersByRequesterID(requesterID)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get transfers by requester ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTransfersByRequesterID OK")
	return transfers, nil
}

func (s *TransferService) UpdateTransfer(id string, transfer Transfer) error {
	s.Logger.Infof("Service UpdateTransfer")

	if err := s.TransferRepo.UpdateTransfer(id, transfer); err != nil {
		s.Logger.WithError(err).Error("failed to update transfer")
		return err
	}
	s.Logger.Infof("Service UpdateTransfer OK")
	return nil
}

