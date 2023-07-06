package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type TransferService struct {
	TransferRepo    *TransferRepository
	ProtocolService *utils.ProtocolService
	FilesTransferService *FilesTransferService
	Logger          *logrus.Logger
}

func NewTransferService(transferRepo *TransferRepository, protocolService *utils.ProtocolService, fileTransferService *FilesTransferService, logger *logrus.Logger) *TransferService {
	return &TransferService{
		TransferRepo:    transferRepo,
		ProtocolService: protocolService,
		FilesTransferService: fileTransferService,
		Logger:          logger,
	}
}

func (s *TransferService) CreateTransfer(req Transfer, filesWithDesc []utils.FileWithDescription) (*Transfer, error) {
	s.Logger.Infof("Service CreateTransfer")

	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("P")
	req.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generating protocol for transfer: %v", err)
		return nil, err
	}
	req.Status = "submitted"
	transfer, err := s.TransferRepo.CreateTransfer(req)
	if err != nil {
		s.Logger.Errorf("error fetching transfer from repository: %v", err)
		return nil, err
	}
	// Check if there are files to save
	if len(filesWithDesc) > 0 {
		// Save the files for this cat
		filesTransfer, err := s.FilesTransferService.SaveTransferFiles(transfer.ID, filesWithDesc)
		if err != nil {
			s.Logger.Errorf("error saving Transfer files: %v", err)
			return nil, err
		}

		transfer.Files = &filesTransfer
	} else {
		s.Logger.Info("No files to save for this Transfer")
	}


	s.Logger.Infof("Service CreateTransfer OK")
	return &transfer, nil
}

func (s *TransferService) GetTransferByID(id string) (*Transfer, error) {
	s.Logger.Infof("Service GetTransferByID")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	transfer, err := s.TransferRepo.GetTransferByID(uint(idUint))
	if err != nil {
		s.Logger.Errorf("error fetching transfer from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetTransferByID OK")
	return &transfer, nil
}

func (s *TransferService) UpdateTransfer(id string, transfer Transfer) error {
	s.Logger.Infof("Service UpdateTransfer")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	if err := s.TransferRepo.UpdateTransfer(uint(idUint), transfer); err != nil {
		s.Logger.WithError(err).Error("failed to update transfer")
		return err
	}
	s.Logger.Infof("Service UpdateTransfer OK")
	return nil
}

func (s *TransferService) UpdateTransferStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateTransferStatus")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	// check if the transfer exists
	if _, err := s.TransferRepo.GetTransferByID(uint(idUint)); err != nil {
		s.Logger.WithError(err).Error("transfer not found")
		return err
	}

	if err := s.TransferRepo.UpdateTransferStatus(uint(idUint), status); err != nil {
		s.Logger.WithError(err).Error("failed to update transfer status")
		return err
	}

	s.Logger.Infof("Service UpdateTransferStatus OK")
	return nil
}



func (s *TransferService) GetAllTransfersByRequesterID(requesterID string) ([]Transfer, error) {
	s.Logger.Infof("Service GetAllTransfersByRequesterID")
	
	idUint, err := strconv.ParseUint(requesterID, 10, 64)
	if err != nil {
		return nil, err
	}

	transfers, err := s.TransferRepo.GetAllTransfersByRequesterID(uint(idUint))
	if err != nil {
		s.Logger.WithError(err).Error("failed to get transfers by requester ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTransfersByRequesterID OK")
	return transfers, nil
}


