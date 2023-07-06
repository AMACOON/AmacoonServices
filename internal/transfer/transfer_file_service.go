package transfer

import (
	"strconv"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type FilesTransferService struct {
	FileService *utils.FilesService
	Logger      *logrus.Logger
	FilesTransferRepo     *FilesTransferRepository
}

func NewFilesTransferService(fileService *utils.FilesService, filesTransferRepo *FilesTransferRepository, logger *logrus.Logger) *FilesTransferService {
	return &FilesTransferService{
		FileService:  fileService,
		FilesTransferRepo: filesTransferRepo,
		Logger:       logger,
	}
}

func (s *FilesTransferService) SaveTransferFiles(TransferID uint, filesWithDesc []utils.FileWithDescription) ([]FilesTransfer, error) {
	s.Logger.Infof("Service SaveTransferFiles")

	// Save the files using the FilesService
	files, err := s.FileService.SaveFiles(strconv.FormatUint(uint64(TransferID), 10), "transfers", filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error save transfers files: %v", err)
		return nil, err
	}

	// Convert the saved files into FilesTransfer and save them
	var filesTransferCreated []FilesTransfer
	for _, file := range files {
		fileTransfer := FilesTransfer{
			TransferID:    TransferID,
			FileData: file,
		}
		created, err := s.FilesTransferRepo.CreateFilesTransfer([]FilesTransfer{fileTransfer})
		if err != nil {
			s.Logger.Errorf("Failed to create file Transfer: %v", err)
			return nil, err
		}
		filesTransferCreated = append(filesTransferCreated, created...)
	}

	s.Logger.Infof("Service SaveTransferFiles OK")
	return filesTransferCreated, nil
}
