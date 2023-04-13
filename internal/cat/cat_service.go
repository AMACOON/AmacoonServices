package cat

import (
	
	"github.com/sirupsen/logrus"
)


type CatService struct {
	CatRepo *CatRepository
	Logger        *logrus.Logger
}

func NewCatService(catRepo *CatRepository, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo: catRepo,
        Logger:       logger,
	}
}

func (s *CatService) GetCatsCompleteByID(id string) (*CatComplete, error) {
    cats, err := s.CatRepo.GetCatCompleteByID(id)
    if err != nil {
        s.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex from repo")
        return nil, err
    }
    return cats, nil
}

// func (s *CatService) GetCatByRegistrationTable(registration string) (*CatTable, error) {
//     cat, err := s.CatRepo.GetCatByRegistrationTable(registration)
//     if err != nil {
//         s.Logger.WithError(err).Errorf("Failed to get cat by registration '%s' from repo", registration)
//         return nil, err
//     }
//     return cat, nil
// }

// func (s *CatService) GetCatsByExhibitorAndSex(idExhibitor int, sex int) ([]Cat, error) {
//     cats, err := s.CatRepo.GetCatsByExhibitorAndSex(idExhibitor, sex)
//     if err != nil {
//         s.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex from repo")
//         return nil, err
//     }
//     return cats, nil
// }

// func (s *CatService) GetCatByRegistration(registration string) (*Cat, error) {
//     cat, err := s.CatRepo.GetCatByRegistration(registration)
//     if err != nil {
//         s.Logger.WithError(err).Errorf("Failed to get cat by registration '%s' from repo", registration)
//         return nil, err
//     }
//     return cat, nil
// }

