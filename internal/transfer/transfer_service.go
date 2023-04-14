package transfer

import (
	"errors"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type TransferService struct {
	TransferRepo    *TransferRepository
	ProtocolService *utils.ProtocolService
	Logger          *logrus.Logger
}

func NewTransferService(transferRepo *TransferRepository, logger *logrus.Logger, protocolService *utils.ProtocolService) *TransferService {
	return &TransferService{
		TransferRepo:    transferRepo,
		ProtocolService: protocolService,
		Logger:          logger,
	}
}

func (s *TransferService) CreateTransfer(transfer Transfer) (Transfer, error) {
	s.Logger.Infof("Service CreateTransfer")

	protocolNumber := s.ProtocolService.GenerateProtocolNumber("T")
	
	transfer.ProtocolNumber = protocolNumber

	transfer, err := s.TransferRepo.CreateTransfer(transfer)
	if err != nil {
		s.Logger.Errorf("error creating transfer in repository: %v", err)
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

func (s *TransferService) GetLitterFilesByID(id string) ([]utils.Files, error) {
	s.Logger.Infof("Service GetTransferFilesByID")

	// check if the litter exists
	if _, err := s.TransferRepo.GetTransferByID(id); err != nil {
		return nil, errors.New("litter not found")
	}

	files, err := s.TransferRepo.GetTransferFilesByID(id)
	if err != nil {
		return nil, err
	}

	s.Logger.Infof("Service GetTransferFilesByID OK")
	return files, nil
}

func (s *TransferService) GetAllTransfersByOwner(requesterID string) ([]Transfer, error) {
	s.Logger.Infof("Service GetAllTransfersByRequester")
	transfers, err := s.TransferRepo.GetAllTransfersByOwner(requesterID)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get transfers by requester ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTransfersByRequester OK")
	return transfers, nil
}

func (s *TransferService) UpdateTransfer(id string,transfer Transfer) error {
	s.Logger.Infof("Service UpdateTransfer")
	if transfer.ID.IsZero() {
		err := errors.New("invalid transfer ID")
		s.Logger.Errorf("error updating transfer: %v", err)
		return err
	}
	if err := s.TransferRepo.UpdateTransfer(id ,transfer); err != nil {
		s.Logger.WithError(err).Error("failed to update transfer")
		return err
	}
	s.Logger.Infof("Service UpdateTransfer OK")
	return nil
}


