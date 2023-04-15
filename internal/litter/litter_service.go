package litter

import (
	"errors"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type LitterService struct {
	LitterRepo      *LitterRepository
	ProtocolService *utils.ProtocolService
	Logger          *logrus.Logger
}

func NewLitterService(litterRepo *LitterRepository, logger *logrus.Logger, protocolService *utils.ProtocolService) *LitterService {
	return &LitterService{
		LitterRepo:      litterRepo,
		ProtocolService: protocolService,
		Logger:          logger,
	}
}

func (s *LitterService) CreateLitter(req LitterRequest) (Litter, error) {
	s.Logger.Infof("Service CreateLitter")

	reqEntidy, err:= ConvertLitterRequestToLitter(req)
	if err != nil {
		s.Logger.Errorf("error convert litterReq to  litter: %v", err)
		return Litter{}, err
	}
	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("L")
	reqEntidy.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generate protocol to  litter: %v", err)
		return Litter{}, err
	}
	
	litter, err := s.LitterRepo.CreateLitter(reqEntidy)
	if err != nil {
		s.Logger.Errorf("error fetching litter from repository: %v", err)
		return Litter{}, err
	}

	s.Logger.Infof("Service CreateLitter OK")
	return litter, nil
}

func (s *LitterService) GetLitterByID(id string) (*Litter, error) {
	s.Logger.Infof("Service GetLitterByID")

	litter, err := s.LitterRepo.GetLitterByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching litter from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetLitterByID OK")
	return &litter, nil
}

func (s *LitterService) UpdateLitterStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateLitterStatus")

	// check if the litter exists
	if _, err := s.LitterRepo.GetLitterByID(id); err != nil {
		return errors.New("litter not found")
	}

	if err := s.LitterRepo.UpdateLitterStatus(id, status); err != nil {
		return err
	}

	s.Logger.Infof("Service UpdateLitterStatus OK")
	return nil
}

func (s *LitterService) AddLitterFiles(id string, files []utils.Files) error {
	s.Logger.Infof("Service AddLitterFiles")

	// check if the litter exists
	if _, err := s.LitterRepo.GetLitterByID(id); err != nil {
		return errors.New("litter not found")
	}

	if err := s.LitterRepo.AddLitterFiles(id, files); err != nil {
		return err
	}

	s.Logger.Infof("Service AddLitterFiles OK")
	return nil
}

// func (s *LitterService) GetLitterFilesByID(id string) ([]utils.Files, error) {
// 	s.Logger.Infof("Service GetLitterFilesByID")

// 	// check if the litter exists
// 	if _, err := s.LitterRepo.GetLitterByID(id); err != nil {
// 		return nil, errors.New("litter not found")
// 	}

// 	files, err := s.LitterRepo.GetLitterFilesByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s.Logger.Infof("Service GetLitterFilesByID OK")
// 	return files, nil
// }

func (s *LitterService) GetAllLittersByOwner(ownerId string) ([]Litter, error) {
	s.Logger.Infof("Service GetLittersByOwnerId")
	litters, err := s.LitterRepo.GetAllLittersByOwner(ownerId)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get litters by owner ID")
		return nil, err
	}
	s.Logger.Infof("Service GetLittersByOwnerId OK")
	return litters, nil
}

func (s *LitterService) UpdateLitter(id string ,litter Litter) error {
	s.Logger.Infof("Service UpdateLitter")

	if err := s.LitterRepo.UpdateLitter(id, litter); err != nil {
		s.Logger.WithError(err).Error("failed to update litter")
		return err
	}
	s.Logger.Infof("Service UpdateLitter OK")
	return nil
}
