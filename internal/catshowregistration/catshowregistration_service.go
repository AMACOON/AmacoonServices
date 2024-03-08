package catshowregistration

import (
	"errors"
	"fmt"

	"encoding/json"
	"log"

	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type CatShowRegistrationService struct {
	Logger            *logrus.Logger
	CatShowCatService *catshowcat.CatShowCatService
	CatService        *cat.CatService
	Repo              *CatShowRegistrationRepository
}

func NewCatShowRegistrationService(logger *logrus.Logger, catShowcatService *catshowcat.CatShowCatService, catService *cat.CatService, repo *CatShowRegistrationRepository) *CatShowRegistrationService {
	return &CatShowRegistrationService{
		Logger:            logger,
		CatShowCatService: catShowcatService,
		CatService:        catService,
		Repo:              repo,
	}
}

func (s *CatShowRegistrationService) CreateCatShowRegistration(registration *Registration, filesWithDesc []utils.FileWithDescription) (*Registration, error) {
	s.Logger.Infof("Service CreateCatShowRegistration")

	// Cria a inscrição através do repositório
	reg, err := s.Repo.CreateCatShowRegistration(registration)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to create registration")
		return nil, err
	}

	// Obter detalhes completos do gato, incluindo títulos e arquivos
	cat, err := s.CatService.GetCatsCompleteByID(fmt.Sprintf("%d", *registration.CatID))
	if err != nil {
		// Primeiro, deleta o registro do cat show se houver algum erro
		deleteErr := s.Repo.DeleteCatShowRegistrationByID(reg.ID)
		if deleteErr != nil {
			// Loga o erro de deleção, se ocorrer
			s.Logger.WithError(deleteErr).Error("Failed to delete cat show registration")
		}

		// Então, loga o erro original e retorna
		s.Logger.WithError(err).Error("Failed to get cat details")
		return nil, err
	}

	// Fazendo o marshal do objeto cat para JSON
	catJSON, err := json.Marshal(cat)
	if err != nil {
		// Se houver erro no marshal, loga o erro
		s.Logger.WithError(err).Error("Failed to marshal cat details to JSON")
		return nil, err
	}

	// Logando o JSON do objeto cat
	s.Logger.Infof("Detalhes completos do gato: %s", string(catJSON))

	// Tranforma CatShowCat
	// Inicializa a variável para os títulos convertidos
	if cat.Titles != nil {
		s.Logger.Infof("Quantidade de títulos de gato: %d", len(*cat.Titles))
	} else {
		s.Logger.Info("Nenhum título de gato disponível")
	}

	convertedTitles, err := transformTitles(cat.Titles)
	if err != nil {
		s.Repo.DeleteCatShowRegistrationByID(reg.ID)
		s.Logger.WithError(err).Error("Failed to convert titles")
		return nil, err
	}

	catShowCat, err := createCatShowCatFromCat(cat, convertedTitles)
	if err != nil {
		s.Repo.DeleteCatShowRegistrationByID(reg.ID)
		s.Logger.WithError(err).Error("Failed to create CatShowCat from Cat")
		return nil, err
	}
	catShowCat.RegistrationID = &reg.ID

	// Cria o registro de CatShowCat
	_, err = s.CatShowCatService.CreateCatShowCat(&catShowCat, filesWithDesc)
	if err != nil {
		// Log do erro
		s.Repo.DeleteCatShowRegistrationByID(reg.ID)
		s.Logger.WithError(err).Error("Failed to create CatShowCat")
		return nil, err
	}

	s.Logger.Infof("Service CreateCatShowRegistration OK")
	return reg, nil
}

func transformTitles(catTitles *[]cat.TitlesCat) ([]catshowcat.TitlesCatShowCat, error) {
	var convertedTitles []catshowcat.TitlesCatShowCat

	if catTitles != nil && len(*catTitles) > 0 {
		log.Printf("Transformando %d títulos de gato.", len(*catTitles))
		for _, titleCat := range *catTitles {
			convertedTitle := catshowcat.TitlesCatShowCat{
				TitleID:      titleCat.TitleID,
				Date:         titleCat.Date,
				FederationID: titleCat.FederationID,
			}
			convertedTitles = append(convertedTitles, convertedTitle)
		}
		log.Println("Títulos de gato transformados com sucesso.")
	} else {
		log.Println("Nenhum título para transformar.")
	}
	return convertedTitles, nil
}

func createCatShowCatFromCat(cat *cat.Cat, convertedTitles []catshowcat.TitlesCatShowCat) (catshowcat.CatShowCat, error) {
	if cat == nil {
		return catshowcat.CatShowCat{}, errors.New("cat is nil")
	}

	return catshowcat.CatShowCat{
		Name:                cat.Name,
		Registration:        cat.Registration,
		RegistrationType:    cat.RegistrationType,
		Microchip:           cat.Microchip,
		Gender:              cat.Gender,
		Birthdate:           cat.Birthdate,
		Neutered:            cat.Neutered,
		Validated:           cat.Validated,
		Observation:         cat.Observation,
		Fifecat:             cat.Fifecat,
		FederationID:        cat.FederationID,
		BreedID:             cat.BreedID,
		ColorID:             cat.ColorID,
		CatteryID:           cat.CatteryID,
		OwnerID:             cat.OwnerID,
		CountryID:           cat.CountryID,
		Titles:              &convertedTitles,
		FatherID:            cat.FatherID,
		FatherName:          cat.FatherName,
		FatherBreedId:       cat.FatherBreedId,
		FatherColorID:       cat.FatherColorID,
		FatherNameManual:    cat.FatherNameManual,
		FatherBreedIDManual: cat.FatherBreedIDManual,
		FatherColorIDManual: cat.FatherColorIDManual,
		MotherID:            cat.MotherID,
		MotherName:          cat.MotherName,
		MotherBreedID:       cat.MotherBreedID,
		MotherColorId:       cat.MotherColorId,
		MotherNameManual:    cat.MotherNameManual,
		MotherBreedIDManual: cat.MotherBreedIDManual,
		MotherColorIDManual: cat.MotherColorIDManual,
		FatherNameTemp:      cat.FatherNameTemp,
		MotherNameTemp:      cat.MotherNameTemp,
	}, nil
}
