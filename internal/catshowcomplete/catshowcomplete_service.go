package catshowcomplete

import (
	"github.com/sirupsen/logrus"
)


// CatShowCompleteService fornece métodos para operações de serviço em CatShowComplete.
type CatShowCompleteService struct {
	Logger *logrus.Logger
	Repo   *CatShowCompleteRepository
	
}

// NewCatShowCompleteService cria uma nova instância de CatShowCompleteService.
func NewCatShowCompleteService(logger *logrus.Logger, repo *CatShowCompleteRepository) *CatShowCompleteService {
	return &CatShowCompleteService{
		Logger: logger,
		Repo:   repo,
		
	}
}

// GetCatShowCompleteByID busca informações completas do cat show por registrationID.
func (s *CatShowCompleteService) GetCatShowCompleteByID(registrationID uint) (*CatShowComplete, error) {
	s.Logger.Infof("Service GetCatShowCompleteByID")
	catShowComplete, err := s.Repo.GetCatShowCompleteByID(registrationID)
	if err != nil {
		s.Logger.Errorf("Failed to get CatShowComplete by ID: %v", err)
		return nil, err
	}
	return catShowComplete, nil
}

// GetCatShowCompleteByCatID busca informações completas do cat show por catID.
func (s *CatShowCompleteService) GetCatShowCompleteByCatID(catID uint) ([]CatShowComplete, error) {
	s.Logger.Infof("Service GetCatShowCompleteByCatID")
	catShowCompletes, err := s.Repo.GetCatShowCompleteByCatID(catID)
	if err != nil {
		s.Logger.Errorf("Failed to get CatShowComplete by CatID: %v", err)
		return nil, err
	}
	return catShowCompletes, nil
}

// GetCatShowCompleteByCatShowIDs busca informações completas do cat show por catShowID e catShowSubID (opcional).
func (s *CatShowCompleteService) GetCatShowCompleteByCatShowIDs(catShowID uint, catShowSubID *uint) ([]CatShowComplete, error) {
	s.Logger.Infof("Service GetCatShowCompleteByCatShowIDs")
	catShowCompletes, err := s.Repo.GetCatShowCompleteByCatShowIDs(catShowID, catShowSubID)
	if err != nil {
		s.Logger.Errorf("Failed to get CatShowComplete by CatShowIDs: %v", err)
		return nil, err
	}
	return catShowCompletes, nil
}
