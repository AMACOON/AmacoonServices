package title

import (
	
	"github.com/sirupsen/logrus"
)

type TitleService struct {
	TitleRepo *TitleRepository
	Logger        *logrus.Logger
}

func NewTitleService(titleRepo *TitleRepository, logger *logrus.Logger) *TitleService {
	return &TitleService{
		TitleRepo: titleRepo,
		Logger:       logger,
	}
}

func (s *TitleService) GetAllTitles() ([]TitleMongo, error) {
	s.Logger.Infof("Service GetAllTitles")
	
	titles, err := s.TitleRepo.GetAllTitles()
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get all countries")
		return nil, err
	}
	s.Logger.Infof("Service GetAllTitles OK")
	return titles, nil
}



