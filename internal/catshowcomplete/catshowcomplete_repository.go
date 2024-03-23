package catshowcomplete

import (
	"sort"

	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
	"github.com/scuba13/AmacoonServices/internal/judge"
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
		//Preload("CatShow.Federation").
		//Preload("CatShow.Club").
		//Preload("CatShow.Country").
		Preload("CatShowSub").
		Preload("Owner").
		//Preload("Owner.Country").
		//Preload("Owner.Clubs").
		Preload("Class").
		Preload("Judge").
		Preload("Judge.Country").
		Preload("CatShowCat").
		//Preload("CatShowCat.Breed").
		//Preload("CatShowCat.Color").
		//Preload("CatShowCat.Federation").
		//Preload("CatShowCat.Country").
		//Preload("CatShowCat.Cattery").
		//Preload("CatShowCat.Cattery.Country").
		//Preload("CatShowCat.Cattery.Owner").
		//Preload("CatShowCat.Owner").
		//Preload("CatShowCat.Owner.Country").
		//Preload("CatShowCat.Owner.Clubs").
		//Preload("CatShowCat.Titles").
		//Preload("CatShowCat.Files").
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

// CatShowYearGroup agrupa os CatShowDetails por ano.
type CatShowYearGroup struct {
	Year     int
	CatShows []CatShowDetail
}

// CatShowDetail representa detalhes de um CatShow, incluindo suas subdivisões e resultados relacionados.
type CatShowDetail struct {
	CatShow     catshow.CatShow
	CatShowSubs []CatShowSubDetail
}

// CatShowSubDetail combina informações de uma subexposição, resultados e detalhes do gato.
type CatShowSubDetail struct {
	CatShowSub catshow.CatShowSub
	Results    catshowresult.CatShowResult
	Judges     judge.Judge
	CatShowCat catshowcat.CatShowCat
}

// Método do repositório para buscar os dados completos por ID do gato.
func (r *CatShowCompleteRepository) GetCatShowCompleteByYear(catID uint) ([]CatShowYearGroup, error) {
	r.Logger.Info("Fetching CatShowComplete by CatID")

	var registrations []catshowregistration.Registration
	if err := r.DB.
		Preload("CatShow").
		Preload("CatShow.Federation").
		Preload("CatShow.Club").
		Preload("CatShow.Country").
		Preload("CatShowSub").
		Preload("Owner").
		Preload("Class").
		Preload("Judge").
		Preload("Judge.Country").
		Preload("CatShowCat").
		Preload("CatShowCat.Breed").
		Preload("CatShowCat.Color").
		Preload("CatShowCat.Owner").
		Where("cat_id = ?", catID).
		Find(&registrations).Error; err != nil {
		r.Logger.Error("Error fetching Registrations by CatID: ", err)
		return nil, err
	}

	// Estrutura para manter os CatShows agrupados por ano.
	yearGroupMap := make(map[int]map[uint]CatShowDetail)

	for _, reg := range registrations {
		year := reg.CatShow.StartDate.Year()
		if yearGroupMap[year] == nil {
			yearGroupMap[year] = make(map[uint]CatShowDetail)
		}

		catShowID := reg.CatShow.ID
		catShowDetail, exists := yearGroupMap[year][catShowID]
		if !exists {
			catShowDetail = CatShowDetail{
				CatShow:     *reg.CatShow,
				CatShowSubs: []CatShowSubDetail{},
			}
		}

		if reg.CatShowSub != nil && reg.CatShowSub.ID != 0 {
			var result catshowresult.CatShowResult
			if err := r.DB.Preload("CatShowResultMatrix").Where("cat_show_sub_id = ?", reg.CatShowSub.ID).First(&result).Error; err == nil {
				catShowSubDetail := CatShowSubDetail{
					CatShowSub: *reg.CatShowSub,
					Results:    result,
					Judges:     *reg.Judge,
					CatShowCat: *reg.CatShowCat,
				}
				catShowDetail.CatShowSubs = append(catShowDetail.CatShowSubs, catShowSubDetail)
			}
		}

		yearGroupMap[year][catShowID] = catShowDetail
	}

	var catShowYearGroups []CatShowYearGroup
	for year, detailsMap := range yearGroupMap {
		var catShows []CatShowDetail
		for _, detail := range detailsMap {
			catShows = append(catShows, detail)
		}
		catShowYearGroups = append(catShowYearGroups, CatShowYearGroup{
			Year:     year,
			CatShows: catShows,
		})
	}
	// Após ter construído catShowYearGroups, ordene-os por ano, decrescentemente.
	sort.Slice(catShowYearGroups, func(i, j int) bool {
		// Retorna verdadeiro se o ano do elemento i for maior que o ano do elemento j
		return catShowYearGroups[i].Year > catShowYearGroups[j].Year
	})

	r.Logger.Info("Successfully fetched CatShowComplete data")
	return catShowYearGroups, nil
}
