package litter

import (
	"github.com/scuba13/AmacoonServices/internal/utils"

	"github.com/sirupsen/logrus"
)

type LitterService struct {
	LitterRepo      *LitterRepository
	FilesRepo       *utils.FilesRepository
	Logger          *logrus.Logger
	LitterConverter *LitterConverter
}

func NewLitterService(litterRepo *LitterRepository, filesRepo *utils.FilesRepository, logger *logrus.Logger, litterConverter *LitterConverter) *LitterService {
	return &LitterService{
		LitterRepo:      litterRepo,
		FilesRepo:       filesRepo,
		Logger:          logger,
		LitterConverter: litterConverter,
	}
}

func (s *LitterService) GetAllLitters() ([]Litter, error) {
	littersDB, err := s.LitterRepo.GetAllLitters()
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get all litters")
		return nil, err
	}

	var litters []Litter
	for _, litterDB := range littersDB {
		kittensDB, err := s.LitterRepo.GetKittensByLitterID(litterDB.ID)
		if err != nil {
			s.Logger.WithError(err).Error("Failed to get kittens by litter ID")
			return nil, err
		}

		filesDB, err := s.FilesRepo.GetFilesByServiceID(litterDB.ID)
		if err != nil {
			s.Logger.WithError(err).Error("Failed to get files by service ID")
			return nil, err
		}

		litter := s.LitterConverter.LitterDBToLitter(&litterDB, kittensDB, filesDB)
		litters = append(litters, *litter)
	}

	return litters, nil
}

func (s *LitterService) GetLitterByID(litterID uint) (*Litter, error) {
	// Call the repository to get the litter and its kittens
	litter, kittens, files, err := s.LitterRepo.GetLitterByID(litterID)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to get litter with ID: %d", litterID)
		return nil, err
	}

	// Transform the models.Litter and []*models.Kitten into a models.LitterData struct
	litterData := s.LitterConverter.LitterDBToLitter(litter, kittens, files)

	return litterData, nil
}


func (s *LitterService) CreateLitter(litter Litter) (uint, string, error) {
    // Transform LitterData into a models.LitterDB struct
    litterDB, kittensDB, filesDB := s.LitterConverter.LitterToLitterDB(litter)

    // Call the repository to create the litter and its kittens
    litterID, protocolNumber, err := s.LitterRepo.CreateLitter(&litterDB, kittensDB, filesDB)
    if err != nil {
        s.Logger.WithError(err).Error("Failed to create litter")
        return 0, "", err
    }

    return litterID, protocolNumber, nil
}

func (s *LitterService) UpdateLitter(litterID uint, litter Litter) error {
    // Transform LitterData into a models.LitterDB struct
    litterDB, kittensDB, filesDB := s.LitterConverter.LitterToLitterDB(litter)

    // Call the repository to update the litter and its kittens
    if err := s.LitterRepo.UpdateLitter(litterID, &litterDB, kittensDB, filesDB); err != nil {
        s.Logger.WithError(err).Error("Failed to update litter")
        return err
    }

    return nil
}

func (s *LitterService) DeleteLitter(litterID uint) error {
	// Call the repository to delete the litter and its kittens
	err := s.LitterRepo.DeleteLitter(litterID)
	if err != nil {
        s.Logger.WithError(err).Error("Failed to delete litter")
		return err
	}
	return nil
}


