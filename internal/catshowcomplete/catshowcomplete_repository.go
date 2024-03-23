package catshowcomplete

import (
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatShowCompleteRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatShowCompleteRepository(db *gorm.DB, logger *logrus.Logger) *CatShowCompleteRepository {
	return &CatShowCompleteRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatShowCompleteRepository) GetCatShowCompleteByID(registrationID uint) (*CatShowComplete, error) {
    r.Logger.Infof("Repository GetCatShowCompleteByID")
    
    // Primeiro, carregamos a entidade Registration e suas associações diretas.
    var registration catshowregistration.Registration
    if err := r.DB.Preload("CatShow").
        Preload("CatShowSub").
        Preload("Owner").
        Preload("Class").
        Preload("Judge").
        First(&registration, "id = ?", registrationID).Error; err != nil {
        r.Logger.Errorf("Error fetching Registration with preloads: %v", err)
        return nil, err
    }

    // Em seguida, buscamos separadamente o CatShowResult associado e sua matriz.
    var result catshowresult.CatShowResult
    if err := r.DB.Where("registration_id = ?", registrationID).
        Preload("CatShowResultMatrix").
        First(&result).Error; err != nil {
        r.Logger.Errorf("Error fetching CatShowResult with CatShowResultMatrix by RegistrationID: %v", err)
        return nil, err
    }

    // Compondo o CatShowComplete com os resultados das consultas anteriores.
    catShowComplete := &CatShowComplete{
        Registration:  registration,
        CatShowResult: result,
    }

    r.Logger.Infof("Repository GetCatShowCompleteByID OK")
    return catShowComplete, nil
}



func (r *CatShowCompleteRepository) GetCatShowCompleteByCatID(catID uint) ([]CatShowComplete, error) {
	r.Logger.Infof("Repository GetCatShowCompleteByCatID")
	var registrations []catshowregistration.Registration
	if err := r.DB.
					Preload("CatShow").
					Preload("CatShow.Federation").
					Preload("CatShow.Club").
					Preload("CatShow.Country").
					Preload("CatShowSub").
					Preload("Owner").
					Preload("Owner.Country").
					Preload("Owner.Clubs").
					Preload("Class").
					Preload("Judge").
					Preload("Judge.Country").
					Preload("CatShowCat").
					Preload("CatShowCat.Breed").
					Preload("CatShowCat.Color").
					Preload("CatShowCat.Federation").
					Preload("CatShowCat.Country").
					Preload("CatShowCat.Cattery").
					Preload("CatShowCat.Cattery.Country").
					Preload("CatShowCat.Cattery.Owner").
					Preload("CatShowCat.Owner").
					Preload("CatShowCat.Owner.Country").
					Preload("CatShowCat.Owner.Clubs").
					Preload("CatShowCat.Titles").
					Preload("CatShowCat.Files").
					Where("cat_id = ?", catID).Find(&registrations).Error; err != nil {
						r.Logger.Errorf("Error fetching Registrations by CatID: %v", err)
						return nil, err
					}

	var catShowCompletes []CatShowComplete
	for _, registration := range registrations {
		var result catshowresult.CatShowResult
		if err := r.DB.Preload("CatShowResultMatrix").Where("registration_id = ?", registration.ID).First(&result).Error; err != nil {
			r.Logger.Errorf("Error fetching CatShowResult by RegistrationID for CatID %d: %v", catID, err)
			continue // Decida se quer continuar com os próximos ou retornar um erro
		}
		catShowCompletes = append(catShowCompletes, CatShowComplete{
			Registration:  registration,
			CatShowResult: result,
		})
	}

	if len(catShowCompletes) == 0 {
		r.Logger.Errorf("No CatShowComplete found for CatID %d", catID)
		return nil, gorm.ErrRecordNotFound
	}
	r.Logger.Infof("Repository GetCatShowCompleteByCatID OK")
	return catShowCompletes, nil
}

func (r *CatShowCompleteRepository) GetCatShowCompleteByCatShowIDs(catShowID uint, catShowSubID *uint) ([]CatShowComplete, error) {
	r.Logger.Infof("Repository GetCatShowCompleteByCatShowIDs")
	var query *gorm.DB

	// Constrói a consulta baseada na presença do CatShowSubID
	if catShowSubID != nil {
		query = r.DB.Where("cat_show_id = ? AND cat_show_sub_id = ?", catShowID, *catShowSubID)
	} else {
		query = r.DB.Where("cat_show_id = ?", catShowID)
	}

	var registrations []catshowregistration.Registration
	if err := query.Find(&registrations).Error; err != nil {
		r.Logger.Errorf("Error fetching Registrations: %v", err)
		return nil, err
	}

	var catShowCompletes []CatShowComplete
	for _, registration := range registrations {
		var result catshowresult.CatShowResult
		if err := r.DB.Where("registration_id = ?", registration.ID).First(&result).Error; err != nil {
			r.Logger.Errorf("Error fetching CatShowResult by RegistrationID: %v", err)
			continue // ou retornar erro, dependendo da sua lógica de erro desejada
		}
		catShowCompletes = append(catShowCompletes, CatShowComplete{
			Registration:  registration,
			CatShowResult: result,
		})
	}

	if len(catShowCompletes) == 0 {
		r.Logger.Errorf("No CatShowComplete found")
		return nil, gorm.ErrRecordNotFound // Ajuste conforme necessário
	}
	r.Logger.Infof("Repository GetCatShowCompleteByCatShowIDs OK")
	return catShowCompletes, nil
}
