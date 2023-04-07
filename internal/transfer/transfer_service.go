package transfer

import (
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/utils"
)

type TransferService struct {
	TransferRepo      *TransferRepository
	FilesRepo         *utils.FilesRepository
	Logger            *logrus.Logger
	TransferConverter *TransferConverter
}

func NewTransferService(transferRepo *TransferRepository, filesRepo *utils.FilesRepository, logger *logrus.Logger, transferConverter *TransferConverter) *TransferService {
	return &TransferService{
		TransferRepo:      transferRepo,
		FilesRepo:         filesRepo,
		Logger:            logger,
		TransferConverter: transferConverter,
	}
}

func (s *TransferService) CreateTransfer(transfer Transfer) (uint, string, error) {
	transferDB, filesDB := s.TransferConverter.TransferToTransferDB(transfer)
	transferID, protocolNumber, err := s.TransferRepo.CreateTransfer(&transferDB, filesDB)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to create Transfer")
		return 0, "", err
	}

	return transferID, protocolNumber, nil
}

func (s *TransferService) GetAllTransfers() ([]*Transfer, error) {
	s.Logger.Info("Getting all transfers")
	transfers, err := s.TransferRepo.GetAlltransfers()
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get all transfers")
		return nil, err
	}

	var transferList []*Transfer

	// Transform each TransferDB and FilesDB into a Transfer struct
	for _, transfer := range transfers {
		files, err := s.FilesRepo.GetFilesByServiceID(transfer.ID)
		if err != nil {
			s.Logger.WithError(err).Errorf("Failed to get files by transfer ID %v", transfer.ID)
			return nil, err
		}

		transferData := s.TransferConverter.TransferDBToTransfer(&transfer, files)
		transferList = append(transferList, transferData)
	}
	return transferList, nil
}

func (s *TransferService) GetTransferByID(transferID uint) (*Transfer, error) {
	s.Logger.Info("Getting transfer with ID ", transferID)

	transferDB, filesDB, err := s.TransferRepo.GetTransferByID(transferID)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get transfer from repository")
		return nil, err
	}

	transfer := s.TransferConverter.TransferDBToTransfer(transferDB, filesDB)

	return transfer, nil
}

func (s *TransferService) UpdateTransfer(id uint, transfer *Transfer) error {
    s.Logger.Infof("Updating transfer with id %d", id)
    transferDB, filesDB := s.TransferConverter.TransferToTransferDB(*transfer)
    if err := s.TransferRepo.UpdateTransfer(id, &transferDB, filesDB); err != nil {
        s.Logger.WithError(err).Error("Failed to update transfer")
        return err
    }
    s.Logger.Infof("Successfully updated transfer with id %d", id)
    return nil
}

func (s *TransferService) DeleteTransfer(id uint) error {
    s.Logger.Infof("Deleting transfer with id %d", id)
    if err := s.TransferRepo.DeleteTransfer(id); err != nil {
        s.Logger.WithError(err).Error("Failed to delete transfer")
        return err
    }
    s.Logger.Infof("Successfully deleted transfer with id %d", id)
    return nil
}



