package catshowyear

import (
	"github.com/sirupsen/logrus"
)


//
type CatShowYearService struct {
	Logger *logrus.Logger
	Repo   *CatShowYearRepository
	
}

func NewCatShowYearService(logger *logrus.Logger, repo *CatShowYearRepository) *CatShowYearService {
	return &CatShowYearService{
		Logger: logger,
		Repo:   repo,
		
	}
}

func (s *CatShowYearService) GetCatShowCompleteByYear(catID uint) ([]CatShowYearGroup, error) {
	s.Logger.Infof("Service GetCatShowCompleteByYear")
	catShowYearGroups, err := s.Repo.GetCatShowCompleteByYear(catID)
	if err != nil {
		s.Logger.Errorf("Failed to get CatShowComplete by year: %v", err)
		return nil, err
	}
	return catShowYearGroups, nil
}
