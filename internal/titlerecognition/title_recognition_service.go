package titlerecognition

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type TitleRecognitionService struct {
	TitleRecognitionRepo *TitleRecognitionRepository
	ProtocolService      *utils.ProtocolService
	Logger               *logrus.Logger
}

func NewTitleRecognitionService(titleRecognitionRepo *TitleRecognitionRepository, logger *logrus.Logger, protocolService *utils.ProtocolService) *TitleRecognitionService {
	return &TitleRecognitionService{
		TitleRecognitionRepo: titleRecognitionRepo,
		ProtocolService:      protocolService,
		Logger:               logger,
	}
}

func (s *TitleRecognitionService) CreateTitleRecognition(req TitleRecognition) (TitleRecognition, error) {
	s.Logger.Infof("Service CreateTitleRecognition")

	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("T")
	req.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generating protocol to titleRecognition: %v", err)
		return TitleRecognition{}, err
	}
	req.Status = "submitted"
	titleRecognition, err := s.TitleRecognitionRepo.CreateTitleRecognition(req)
	if err != nil {
		s.Logger.Errorf("error fetching titleRecognition from repository: %v", err)
		return TitleRecognition{}, err
	}

	s.Logger.Infof("Service CreateTitleRecognition OK")
	return titleRecognition, nil
}

func (s *TitleRecognitionService) GetTitleRecognitionByID(id string) (*TitleRecognition, error) {
	s.Logger.Infof("Service GetTitleRecognitionByID")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	titleRecognition, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(uint(idUint))
	if err != nil {
		s.Logger.Errorf("error fetching titleRecognition from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetTitleRecognitionByID OK")
	return &titleRecognition, nil
}

func (s *TitleRecognitionService) UpdateTitleRecognition(id string, titleRecognition TitleRecognition) error {
	s.Logger.Infof("Service UpdateTitleRecognition")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	if err := s.TitleRecognitionRepo.UpdateTitleRecognition(uint(idUint), titleRecognition); err != nil {
		s.Logger.WithError(err).Error("failed to update titleRecognition")
		return err
	}
	s.Logger.Infof("Service UpdateTitleRecognition OK")
	return nil
}

func (s *TitleRecognitionService) UpdateTitleRecognitionStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateTitleRecognitionStatus")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	if _, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(uint(idUint)); err != nil {
		s.Logger.WithError(err).Error("titleRecognition not found")
		return err
	}

	if err := s.TitleRecognitionRepo.UpdateTitleRecognitionStatus(uint(idUint), status); err != nil {
		return err
	}

	s.Logger.Infof("Service UpdateTitleRecognitionStatus OK")
	return nil
}


func (s *TitleRecognitionService) GetAllTitleRecognitionsByRequesterID(requesterID string) ([]TitleRecognition, error) {
	s.Logger.Infof("Service GetAllTitleRecognitionsByRequesterID")
	
	titleRecognitions, err := s.TitleRecognitionRepo.GetAllTitleRecognitionByRequesterID(requesterID)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get titleRecognitions by owner ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTitleRecognitionsByRequesterID OK")
	return titleRecognitions, nil
}


