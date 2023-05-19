package litter

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type LitterService struct {
	LitterRepo      *LitterRepository
	ProtocolService *utils.ProtocolService
	FilesLitterService *FilesLitterService
	Logger          *logrus.Logger
	
}

func NewLitterService(litterRepo *LitterRepository, fileLitterService *FilesLitterService , protocolService *utils.ProtocolService, logger *logrus.Logger) *LitterService {
	return &LitterService{
		LitterRepo:      litterRepo,
		FilesLitterService: fileLitterService,
		ProtocolService: protocolService,
		Logger:          logger,
		
	}
}

func (s *LitterService) CreateLitter(req Litter, filesWithDesc []utils.FileWithDescription) (*Litter, error) {
	s.Logger.Infof("Service CreateLitter")

	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("L")
	req.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generate protocol to  litter: %v", err)
		return nil, err
	}
	req.Status = "submitted"
	litter, err := s.LitterRepo.CreateLitter(req)
	if err != nil {
		s.Logger.Errorf("error create litter from repository: %v", err)
		return nil, err
	}

	// Save the files for this cat
	filesLitter, err := s.FilesLitterService.SaveLitterFiles(litter.ID, filesWithDesc)
	if err != nil {
		s.Logger.Errorf("error saving cat files: %v", err)
		return nil, err
	}

	litter.Files = &filesLitter


	s.Logger.Infof("Service CreateLitter OK")
	return &litter, nil
}

func (s *LitterService) GetLitterByID(id string) (*Litter, error) {
	s.Logger.Infof("Service GetLitterByID")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	litter, err := s.LitterRepo.GetLitterByID(uint(idUint))
	if err != nil {
		s.Logger.Errorf("error fetching litter from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetLitterByID OK")
	return &litter, nil
}

func (s *LitterService) UpdateLitter(id string, litter Litter) error {
	s.Logger.Infof("Service UpdateLitter")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	if err := s.LitterRepo.UpdateLitter(uint(idUint), litter); err != nil {
		s.Logger.WithError(err).Error("failed to update litter status")
		return err
	}
	s.Logger.Infof("Service UpdateLitter OK")
	return nil
}

func (s *LitterService) UpdateLitterStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateLitterStatus")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	// check if the litter exists
	if _, err := s.LitterRepo.GetLitterByID(uint(idUint)); err != nil {
		s.Logger.WithError(err).Error("litter not found")
		return err
	}

	if err := s.LitterRepo.UpdateLitterStatus(uint(idUint), status); err != nil {
		s.Logger.WithError(err).Error("failed to update litter")
		return err
	}

	s.Logger.Infof("Service UpdateLitterStatus OK")
	return nil
}

func (s *LitterService) GetAllLittersByRequesterID(requesterID string) ([]Litter, error) {
	s.Logger.Infof("Service GetAllLittersByRequesterID")

	idUint, err := strconv.ParseUint(requesterID, 10, 64)
	if err != nil {
		return nil, err
	}
	litters, err := s.LitterRepo.GetAllLittersByRequesterID(uint(idUint))
	if err != nil {
		s.Logger.WithError(err).Error("failed to get litters by owner ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllLittersByRequesterID OK")
	return litters, nil
}

func (s *LitterService) DeleteLitter(id string) error {
	s.Logger.Infof("Service DeleteLitter")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	err = s.LitterRepo.DeleteLitter(uint(idUint))
	if err != nil {
		s.Logger.Errorf("error delete litter from repository: %v", err)
		return err
	}

	s.Logger.Infof("Service DeleteLitter OK")
	return nil
}


