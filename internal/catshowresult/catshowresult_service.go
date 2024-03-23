package catshowresult

import (
	"github.com/sirupsen/logrus"
)

type CatShowResultService struct {
	Logger *logrus.Logger
	Repo   *CatShowResultRepository
}

func NewCatShowResultService(logger *logrus.Logger, repo *CatShowResultRepository) *CatShowResultService {
	return &CatShowResultService{
		Logger: logger,
		Repo:   repo,
	}
}

func (s *CatShowResultService) CreateCatShowResult(result *CatShowResult) (*CatShowResult, error) {
	s.Logger.Info("Service CreateCatShowResult ")
	createdResult, err := s.Repo.Create(result)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to create CatShowResult")
		return nil, err
	}
	s.Logger.Info("Service CreateCatShowResult OK")
	return createdResult, nil
}

func (s *CatShowResultService) GetCatShowResultByID(id uint) (*CatShowResult, error) {
	s.Logger.Infof("Service GetCatShowResultByID")
	result, err := s.Repo.GetById(id)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to fetch CatShowResult with ID: %d", id)
		return nil, err
	}
	s.Logger.Infof("Service GetCatShowResultByID OK")
	return result, nil
}

func (s *CatShowResultService) GetCatShowResultByRegistrationID(registrationID uint) (*CatShowResult, error) {
	s.Logger.Infof("Service GetCatShowResultByRegistrationID")
	result, err := s.Repo.GetByRegistrationID(registrationID)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to fetch CatShowResult with RegistrationID: %d", registrationID)
		return nil, err
	}
	s.Logger.Infof("Service GetCatShowResultByRegistrationID OK")
	return result, nil
}

func (s *CatShowResultService) UpdateCatShowResult(id uint, result *CatShowResult) error {
	s.Logger.Infof("Service UpdateCatShowResult")
	err := s.Repo.Update(id, result)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to update CatShowResult")
		return err
	}
	s.Logger.Infof("Service UpdateCatShowResult OK")
	return nil
}

func (s *CatShowResultService) DeleteCatShowResult(id uint) error {
	s.Logger.Infof("Service DeleteCatShowResult")
	err := s.Repo.Delete(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to delete CatShowResult")
		return err
	}
	s.Logger.Infof("Service DeleteCatShowResult OK")
	return nil
}
