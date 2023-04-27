package titlerecognition

import (
	"errors"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
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

func (s *TitleRecognitionService) CreateTitleRecognition(req TitleRecognitionRequest) (TitleRecognitionMongo, error) {
	s.Logger.Infof("Service CreateTitleRecognition")

	reqEntity, err := ConvertTitleRecognitionRequestToTitleRecognitionMongo(req)
	if err != nil {
		s.Logger.Errorf("error converting titleRecognitionReq to titleRecognition: %v", err)
		return TitleRecognitionMongo{}, err
	}
	protocolNumber, err := s.ProtocolService.GenerateUniqueProtocolNumber("T")
	reqEntity.ProtocolNumber = protocolNumber
	if err != nil {
		s.Logger.Errorf("error generating protocol to titleRecognition: %v", err)
		return TitleRecognitionMongo{}, err
	}
	reqEntity.Status = "submitted"
	titleRecognition, err := s.TitleRecognitionRepo.CreateTitleRecognition(reqEntity)
	if err != nil {
		s.Logger.Errorf("error fetching titleRecognition from repository: %v", err)
		return TitleRecognitionMongo{}, err
	}

	s.Logger.Infof("Service CreateTitleRecognition OK")
	return titleRecognition, nil
}

func (s *TitleRecognitionService) GetTitleRecognitionByID(id string) (*TitleRecognitionMongo, error) {
	s.Logger.Infof("Service GetTitleRecognitionByID")

	titleRecognition, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching titleRecognition from repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service GetTitleRecognitionByID OK")
	return &titleRecognition, nil
}

func (s *TitleRecognitionService) UpdateTitleRecognitionStatus(id string, status string) error {
	s.Logger.Infof("Service UpdateTitleRecognitionStatus")

	if _, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(id); err != nil {
		return errors.New("titleRecognition not found")
	}

	if err := s.TitleRecognitionRepo.UpdateTitleRecognitionStatus(id, status); err != nil {
		return err
	}

	s.Logger.Infof("Service UpdateTitleRecognitionStatus OK")
	return nil
}

func (s *TitleRecognitionService) AddTitleRecognitionFiles(id string, files []utils.Files) error {
	s.Logger.Infof("Service AddTitleRecognitionFiles")

	if _, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(id); err != nil {
		return errors.New("titleRecognition not found")
	}

	if err := s.TitleRecognitionRepo.AddTitleRecognitionFiles(id, files); err != nil {
		return err
	}

	s.Logger.Infof("Service AddTitleRecognitionFiles OK")
	return nil
}

func (s *TitleRecognitionService) GetTitleRecognitionFilesByID(id string) ([]utils.Files, error) {
	s.Logger.Infof("Service GetTitleRecognitionFilesByID")

	if _, err := s.TitleRecognitionRepo.GetTitleRecognitionByID(id); err != nil {
		return nil, errors.New("titleRecognition not found")
	}

	files, err := s.TitleRecognitionRepo.GetTitleRecognitionFilesByID(id)
	if err != nil {
		return nil, err
	}

	s.Logger.Infof("Service GetTitleRecognitionFilesByID OK")
	return files, nil
}

func (s *TitleRecognitionService) GetAllTitleRecognitionsByRequesterID(requesterID string) ([]TitleRecognitionMongo, error) {
	s.Logger.Infof("Service GetAllTitleRecognitionsByRequesterID")
	titleRecognitions, err := s.TitleRecognitionRepo.GetAllTitleRecognitionByRequesterID(requesterID)
	if err != nil {
		s.Logger.WithError(err).Error("failed to get titleRecognitions by owner ID")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTitleRecognitionsByRequesterID OK")
	return titleRecognitions, nil
}

func (s *TitleRecognitionService) UpdateTitleRecognition(id string, titleRecognition TitleRecognitionMongo) error {
	s.Logger.Infof("Service UpdateTitleRecognition")

	if err := s.TitleRecognitionRepo.UpdateTitleRecognition(id, titleRecognition); err != nil {
		s.Logger.WithError(err).Error("failed to update titleRecognition")
		return err
	}
	s.Logger.Infof("Service UpdateTitleRecognition OK")
	return nil
}
