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

func (r *ColorService) UpdateColor(id string, updatedColor *Color) error {
	r.Logger.Infof("UpdateColor")

	err := r.ColorRepo.UpdateColor(id, updatedColor)
	if err != nil {
		r.Logger.WithError(err).Error("Failed to UpdateColor color from repo")
		return err
	}
	r.Logger.Infof("UpdateColor OK")
	return nil

}

func (r *ColorService) GetColorById(id string) (*Color, error) {
	r.Logger.Infof("GetColorById")

	color, err := r.ColorRepo.GetColorById(id)
	if err != nil {
		r.Logger.WithError(err).Error("Failed to GetColorById color from repo")
		return nil, err
	}
	r.Logger.Infof("GetColorById OK")
	return color, nil
}


