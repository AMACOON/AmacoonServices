package cattery

import (
	
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"errors"
)

type CatteryService struct {
	CatteryRepo *CatteryRepository
	CatteryFileService *CatteryFileService
	Logger        *logrus.Logger
}

func NewCatteryService(catteryRepo *CatteryRepository, catteryFileService *CatteryFileService, logger *logrus.Logger) *CatteryService {
	return &CatteryService{
		CatteryRepo: catteryRepo,
		CatteryFileService: catteryFileService,
        Logger:       logger,
	}
}

func (s *CatteryService) GetAllCatteries() ([]CatteryInfo, error) {
    s.Logger.Infof("Service GetAllCatteries")
    catteries, err := s.CatteryRepo.GetAllCatteries()
    if err != nil {
        s.Logger.WithError(err).Error("failed to get all catteries")
        return nil, err
    }
    s.Logger.Infof("Service GetAllCatteries OK")
    return catteries, nil
}

func (s *CatteryService) GetCatteryByID(id string) (*Cattery, error) {
	s.Logger.Infof("Service GetCatteryByID")
    cattery, err := s.CatteryRepo.GetCatteryByID(id)
	if err != nil {
		s.Logger.Errorf("error fetching cattery from repository: %v", err)
		return nil, err
	}
    s.Logger.Infof("Service GetCatteryByID OK")
	return cattery, nil
}

func (s *CatteryService) CreateCattery(req *Cattery, filesWithDesc []utils.FileWithDescription) (*Cattery, error) {
	s.Logger.Infof("Service CreateCattery")

	// Verify if a cattery with the same name already exists
	existingCattery, err := s.CatteryRepo.ValidateCattery(req.Name)
	if err != nil {
		s.Logger.Errorf("error finding existing cattery from repository: %v", err)
		return nil, err
	}
	if existingCattery != nil {
		s.Logger.Info("A cattery with the same name already exists '%s'", existingCattery.Name)
		return nil, errors.New("a cattery with the same name already exists")
	}

	cattery, err := s.CatteryRepo.CreateCattery(req)
	if err != nil {
		s.Logger.Errorf("error creating cattery from repository: %v", err)
		return nil, err
	}

	// Check if there are files to save
	if len(filesWithDesc) > 0 {
		// Save the files for this cattery
		filesCattery, err := s.CatteryFileService.SaveCatteryFiles(cattery.ID, filesWithDesc)
		if err != nil {
			s.Logger.Errorf("error saving cattery files: %v", err)
			return nil, err
		}
		cattery.Files = filesCattery
	} else {
		s.Logger.Info("No files to save for this cattery")
	}

	s.Logger.Infof("Service CreateCattery OK")
	return cattery, nil
}


func (s *CatteryService) UpdateCattery(id string, cattery *Cattery) (*Cattery, error) {
	s.Logger.Infof("Service UpdateCattery")
	cattery, err := s.CatteryRepo.UpdateCattery(id, cattery)
	if err != nil {
		s.Logger.Errorf("error updating cattery: %v", err)
		return nil, err
	}
	s.Logger.Infof("Service UpdateCattery OK")
	return cattery, nil
}

