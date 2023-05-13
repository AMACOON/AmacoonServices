package catservice

import (
	"github.com/sirupsen/logrus"
)

type CatServiveService struct {
	CatServiceRepo *CatServiceRepository
	Logger  *logrus.Logger
}

func NewCatServiceService(catServiceRepo *CatServiceRepository, logger *logrus.Logger) *CatServiveService {
	return &CatServiveService{
		CatServiceRepo: catServiceRepo,
		Logger:  logger,
	}
}


func (s *CatServiveService) GetCatServiceByID(id string) (*CatServiceData, error) {
	s.Logger.Infof("Service GetCatServiceByID")
	cats, err := s.CatServiceRepo.GetCatServiceByID(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Id from repo")
		return nil, err
	}

	s.Logger.Infof("Service GetCatServiceByID OK")
	return cats, nil
}

func (s *CatServiveService) GetAllCatsServiceByOwnerAndGender(ownerID string, gender string) ([]CatServiceData, error) {
	s.Logger.Infof("Service GetAllCatsServiceByOwnerAndGender")
	cats, err := s.CatServiceRepo.GetAllCatsServiceByOwnerAndGender(ownerID, gender)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner and Gender from repo")
		return nil, err
	}
	
	s.Logger.Infof("Service GetAllCatsServiceByOwnerAndGender OK")
	return cats, nil
}

func (s *CatServiveService) GetAllCatsServiceByOwner(ownerID string) ([]CatServiceData, error) {
	s.Logger.Infof("Service GetAllCatsServiceByOwner")
	cat, err := s.CatServiceRepo.GetAllCatsServiceByOwner(ownerID)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to get cat by Owner '%s' from repo", ownerID)
		return nil, err
	}
	s.Logger.Infof("Service GetAllCatsServiceByOwner OK")
	return cat, nil
}

func (s *CatServiveService) GetCatServiceByRegistration(registration string) (*CatServiceData, error) {
	s.Logger.Infof("Service GetCatServiceByRegistration")
	cat, err := s.CatServiceRepo.GetCatServiceByRegistration(registration)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return nil, err
	}
	s.Logger.Infof("Service GetAllByOwner OK")
	return cat, nil
}

