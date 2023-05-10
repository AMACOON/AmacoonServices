package color

import (
	
	"github.com/sirupsen/logrus"
)

type ColorService struct {
    ColorRepo *ColorRepository
	Logger        *logrus.Logger
}

func NewColorService(colorRepo *ColorRepository, logger *logrus.Logger) *ColorService {
    return &ColorService{
		ColorRepo: colorRepo,
		Logger:       logger,
	}
}

func (s *ColorService) GetAllColorsByBreed(breedCode string) ([]Color, error) {
	s.Logger.Infof("GetAllColorsByBreed")
	colors, err := s.ColorRepo.GetAllColorsByBreed(breedCode)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to get all colors by breed with ID %s", breedCode)
		return nil, err
	}
	s.Logger.Infof("GetAllColorsByBreed OK")
	return colors, nil
}


