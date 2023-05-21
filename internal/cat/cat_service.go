package cat

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CatService struct {
	CatRepo        *CatRepository
	CatFileService *CatFileService
	Logger         *logrus.Logger
}

func NewCatService(catRepo *CatRepository, catFileService *CatFileService, logger *logrus.Logger) *CatService {
	return &CatService{
		CatRepo:        catRepo,
		CatFileService: catFileService,
		Logger:         logger,
	}
}

func (s *CatService) CreateCat(req *Cat, filesWithDesc []utils.FileWithDescription) (*Cat, error) {
	s.Logger.Infof("Service CreateCat")

	req.Validated = false
	cat, err := s.CatRepo.CreateCat(req)
	if err != nil {
		s.Logger.Errorf("error creating cat from repository: %v", err)
		return nil, err
	}

	// Check if there are files to save
	if len(filesWithDesc) > 0 {
		// Save the files for this cat
		filesCat, err := s.CatFileService.SaveCatFiles(cat.ID, filesWithDesc)
		if err != nil {
			s.Logger.Errorf("error saving cat files: %v", err)
			return nil, err
		}
		cat.Files = &filesCat
	}

	s.Logger.Infof("Service CreateCat OK")
	return cat, nil
}

func (s *CatService) GetCatsCompleteByID(id string) (*Cat, error) {
	s.Logger.Infof("Service GetCatsCompleteByID")
	cats, err := s.CatRepo.GetCatCompleteByID(id)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Id from repo")
		return nil, err
	}

	s.Logger.Infof("Service GetCatsCompleteByID OK")
	return cats, nil
}

func (s *CatService) GetCatsByOwner(ownerID string) ([]CatInfo, error) {
	s.Logger.Infof("Service GetCatsByOwner")

	cats, err := s.CatRepo.GetCatsByOwner(ownerID)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return nil, err
	}
	s.Logger.Infof("Service GetCatsByOwner OK")
	return cats, nil
}

func (s *CatService) UpdateNeuteredStatus(catID string, neutered string) error {
	s.Logger.Infof("Service UpdateNeuteredStatus")

	neuteredBool, err := strconv.ParseBool(neutered)
	if err != nil {
		s.Logger.WithError(err).Errorf("Failed to parse neutered status: %s", neutered)
		return  nil
	}

	err = s.CatRepo.UpdateNeuteredStatus(catID, neuteredBool)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to get cats by Owner from repo")
		return err
	}
	s.Logger.Infof("Service UpdateNeuteredStatus OK")
	return nil
}

func GetFullName(cat *Cat, tipo string) string {
	prefix := ""
	suffix := ""
	wwYears := make([]string, 0)

	for _, titleCat := range *cat.Titles {
		title := titleCat.Titles
		if title.Type == "Championship/Premiorship Titles" {
			prefix += title.Code + " "
		} else if title.Type == "Winner Titles" {
			if title.Code == "WW" {
				wwYears = append(wwYears, titleCat.Date.Format("06"))
			} else {
				prefix += titleCat.Date.Format("06") + " " + title.Code + " "
			}
		} else if title.Type == "Merit Titles" {
			suffix += " " + title.Code
		}
	}

	if len(wwYears) > 0 {
		prefix += "WW'" + strings.Join(wwYears, "'") + " "
	}

	if tipo == "" {
		nomeDoGato := strings.ReplaceAll(cat.Name, "'", "&#39;")
		nomeDoGato = cases.Title(language.English).String(nomeDoGato)

		return prefix + cat.Country.Code + (func() string {
			if cat.Country.Code != "" {
				return "* "
			}
			return ""
		}()) + nomeDoGato + suffix
	}

	if tipo == "1" {
		return prefix
	}

	if tipo == "2" {
		return suffix
	}

	return ""
}



