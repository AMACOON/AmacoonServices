package catshow

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

type CatShowService struct {
	CatShowRepo *CatShowRepository
	Logger      *logrus.Logger
}

func NewCatShowService(catShowRepo *CatShowRepository, logger *logrus.Logger) *CatShowService {
	return &CatShowService{
		CatShowRepo: catShowRepo,
		Logger:      logger,
	}
}

func (s *CatShowService) CreateCatShow(catShow *CatShow) (*CatShow, error) {
	s.Logger.Infof("Service CreateCatShow")
	newCatShow, err := s.CatShowRepo.CreateCatShow(catShow)
	if err != nil {
		s.Logger.Errorf("Error creating CatShow: %v", err)
		return nil, err
	}
	s.Logger.Infof("Service CreateCatShow OK")
	return newCatShow, nil
}

func (s *CatShowService) GetCatShowByID(id string) (*CatShow, error) {
	s.Logger.Infof("Service GetCatShowByID")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		s.Logger.Errorf("Failed to convert ID to uint: %v", err)
		return nil, err
	}
	catShow, err := s.CatShowRepo.GetCatShowByID(uint(uintID))
	if err != nil {
		s.Logger.Errorf("Failed to get CatShow by ID from repo: %v", err)
		return nil, err
	}
	s.Logger.Infof("Service GetCatShowByID OK")
	return catShow, nil
}

func (s *CatShowService) UpdateCatShow(id string, catShow *CatShow) error {
	s.Logger.Infof("Service UpdateCatShow")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		s.Logger.Errorf("Failed to convert ID to uint: %v", err)
		return err
	}
	if err := s.CatShowRepo.UpdateCatShow(uint(uintID), catShow); err != nil {
		s.Logger.Errorf("Failed to update CatShow: %v", err)
		return err
	}
	s.Logger.Infof("Service UpdateCatShow OK")
	return nil
}
