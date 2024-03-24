package catshowyear

import (
	"sort"

	"github.com/scuba13/AmacoonServices/internal/catshowresult"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type CatShowYearRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatShowYearRepository(db *gorm.DB, logger *logrus.Logger) *CatShowYearRepository {
	return &CatShowYearRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatShowYearRepository) GetCatShowCompleteByYear(catID uint) ([]CatShowYearGroup, error) {
	r.Logger.Info("Fetching CatShowComplete by CatID")

	var results []catshowresult.CatShowResult
	if err := r.DB.
		Preload("CatShow").
		Preload("CatShowSub").
		Preload("CatShowResultMatrix").
		Preload("Registration.CatShowCat").
		Preload("Registration.CatShowCat.Breed").
		Preload("Registration.CatShowCat.Color").
		Preload("Registration.CatShowCat.Country").
		Preload("Registration.CatShowCat.Titles.Titles").
		Preload("Registration.Owner").
		Preload("Registration.Class").
		Preload("Registration.Judge").
		Where("cat_id = ?", catID).
		Find(&results).Error; err != nil {
		r.Logger.Error("Error fetching CatShowResults by CatID: ", err)
		return nil, err
	}

	// Mapeamento para agrupar os detalhes dos shows por ano
	catShowDetailsByYear := make(map[int][]CatShowDetail)

	for _, result := range results {
		if result.CatShow == nil || result.CatShowSub == nil {
			continue // Pula resultados sem informações completas do show ou da sub
		}

		year := result.CatShow.StartDate.Year()

		// Verifica se o ano já existe no mapeamento
		if _, ok := catShowDetailsByYear[year]; !ok {
			catShowDetailsByYear[year] = []CatShowDetail{} // Inicializa o slice se ainda não existir
		}

		// Constrói o CatShowSubDetail
		catShowSubDetail := CatShowSubDetail{
			CatShowSubDescription: result.CatShowSub.Description,
			CatShowSubCatShowDate: result.CatShowSub.CatShowDate, // Ajustado para StartDate
			ResultsDescription:    result.CatShowResultMatrix.Description,
			ResultScore:           result.CatShowResultMatrix.Score,
			ClassDescription:      result.Registration.Class.Description,
			ClassName:             result.Registration.Class.Name,
			ClassCode:             result.Registration.Class.Code,
			JudgesName:            result.Registration.Judge.Name,
			CatShowCatNameFull:    "",
			CatShowCatBreedCode:   result.Registration.CatShowCat.Breed.BreedCode,
			CatShowCatEmsCode:     result.Registration.CatShowCat.Color.EmsCode,
			CatShowCatGender:      result.Registration.CatShowCat.Gender,
			CatShowCatBirthdate:   result.Registration.CatShowCat.Birthdate,
			OwnerName:             result.Registration.Owner.Name,
		}

		catShowSubDetail.CatShowCatNameFull = GetFullName(result.Registration.CatShowCat)

		// Adiciona detalhes do show ao mapeamento, considerando o agrupamento correto
		found := false
		for i, detail := range catShowDetailsByYear[year] {
			if detail.CatShowDescription == result.CatShow.Description &&
				detail.CatShowCity == result.CatShow.City {
				// Show já listado, adiciona apenas a sub nova
				catShowDetailsByYear[year][i].CatShowSubs = append(detail.CatShowSubs, catShowSubDetail)
				found = true
				break
			}
		}

		if !found {
			// Novo CatShowDetail para o ano
			catShowDetailsByYear[year] = append(catShowDetailsByYear[year], CatShowDetail{
				CatShowYear:        year,
				CatShowDescription: result.CatShow.Description,
				CatShowCity:        result.CatShow.City,
				CatShowSubs:        []CatShowSubDetail{catShowSubDetail},
			})
		}
	}

	// Constrói a lista final de CatShowYearGroup a partir do mapeamento
	var catShowYearGroups []CatShowYearGroup
	for year, details := range catShowDetailsByYear {
		catShowYearGroups = append(catShowYearGroups, CatShowYearGroup{
			Year:     year, // Agora utilizando a variável 'year' para definir o ano do grupo
			CatShows: details,
		})
	}

	// Ordena os grupos de anos
	sort.Slice(catShowYearGroups, func(i, j int) bool {
		return catShowYearGroups[i].Year > catShowYearGroups[j].Year
	})

	r.Logger.Info("Successfully fetched CatShowComplete data")
	return catShowYearGroups, nil
}


func GetFullName(catShowCat *catshowcat.CatShowCat) string {
    if catShowCat == nil {
        return "" // Retorna uma string vazia se cat for nil
    }

    prefix := ""
    suffix := ""
    wwYears := make([]string, 0)

    // Verifica se cat.Titles é nil antes de iterar
    if catShowCat.Titles != nil {
        for _, titleCat := range *catShowCat.Titles {
            title := titleCat.Titles
            if title != nil { // Adiciona verificação para title não ser nil
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
        }
    }

    if len(wwYears) > 0 {
        prefix += "WW'" + strings.Join(wwYears, "'") + " "
    }

    nomeDoGato := strings.ReplaceAll(catShowCat.Name, "'", "&#39;")
    // Supondo que cases.Title e language.English estejam corretamente importados e utilizados
    nomeDoGato = cases.Title(language.English).String(nomeDoGato)

    countryPrefix := ""
    if catShowCat.Country != nil && catShowCat.Country.Code != "" { // Verifica se cat.Country não é nil antes de acessar Code
        countryPrefix = catShowCat.Country.Code + "* "
    }

    return prefix + countryPrefix + nomeDoGato + suffix
}
