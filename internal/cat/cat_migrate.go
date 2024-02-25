package cat

import (
	"log"
	//"github.com/scuba13/AmacoonServices/internal/utils"

	"gorm.io/gorm"
	"errors"
	"fmt"
)

const batchSize = 200

func MigrateCats(dbOld *gorm.DB, dbNew *gorm.DB) {
	var catMigrations []CatTable
	var err error
	var offset int

	newCats := []string{}
	noTitleCats := []string{}

	for {
		catMigrations, err = GetCatsMigrate(dbOld, offset, batchSize)
		if err != nil {
			log.Printf("Error getting cats to migrate: %v\n", err)
		}

		if len(catMigrations) == 0 {
			break
		}

		log.Printf("Migrating cats %d - %d\n", offset, offset+len(catMigrations)-1)

		for i, cat := range catMigrations {
			log.Printf("Migrating cat %d: %s\n", i, cat.Name)

			var existingCat Cat
			result := dbNew.Where("name = ? AND registration = ?", cat.Name, cat.Registration).First(&existingCat)

			if result.Error == nil {
				log.Printf("Cat already exists: %s", cat.Name)
				continue
			}

			neutered := cat.Neutered == "s"
			validated := cat.Validated == "s"
			fifecat := cat.FifeCat == "s"

			titles := []string{}
			if cat.AdultTitle != "0" && cat.AdultTitle != "" {
				titles = append(titles, cat.AdultTitle)
			}

			if cat.NeuterTitle != "0" && cat.NeuterTitle != "" {
				titles = append(titles, cat.NeuterTitle)
			}

			if cat.WW == "1" {
				titles = append(titles, "WW")
			}
			if cat.SW == "1" {
				titles = append(titles, "SW")
			}
			if cat.NW == "1" {
				titles = append(titles, "NW")
			}
			if cat.JW == "1" {
				titles = append(titles, "JW")
			}
			if cat.DVM == "1" {
				titles = append(titles, "DVM")
			}
			if cat.DSM == "1" {
				titles = append(titles, "DSM")
			}
			if cat.DM == "1" {
				titles = append(titles, "DM")
			}
			
			
			
			
			federationID, err := getFederationID(dbNew, cat.FedName)
			if err != nil {
				log.Printf("Error getting federation ID for cat %d: %v\n", i, err)
				continue
			}

			breedID, err := getBreedID(dbNew, cat.BreedName)
			if err != nil {
				log.Printf("Error getting getBreedID for cat %d: %v\n", i, err)
				continue
			}

			colorID, err := getColorID(dbNew, cat.EmsCode, cat.BreedID)
			if err != nil {
				log.Printf("Error getting getColorID for cat %d: %v\n", i, err)
				continue
			}

			countryID, err := findCountryIdByCode(dbNew, cat.Country)
			if err != nil {
				log.Printf("Error getting findCountryIDByCode for cat %d: %v\n", i, err)
				continue
			}

			catteryID, err := getCatteryID(dbNew, cat.BreederName)
			if err != nil {
				log.Printf("Error getting getCatteryID for cat %d: %v\n", i, err)
				continue
			}

			ownerID, err := getOwnerID(dbNew, cat.OwnerName)
			if err != nil {
				log.Printf("Error getting getOwnerID for cat %d: %v\n", i, err)
				continue
			}

			sexString := ""
			if cat.Sex == "1" {
				sexString = "male"
			} else if cat.Sex == "2" {
				sexString = "female"
			}

		

			catGorm := Cat{
				Name:             cat.Name,
				Registration:     cat.Registration,
				RegistrationType: cat.RegType,
				Microchip:        cat.Microchip,
				Gender:           sexString,
				Birthdate:        cat.BirthDate,
				Neutered:         &neutered,
				Validated:        validated,
				Observation:      "",
				Fifecat:          fifecat,
				FederationID:     federationID,
				BreedID:          uintPtr(breedID),
				ColorID:          uintPtr(colorID),
				FatherID:         nil,
				MotherID:         nil,
				CatteryID:        catteryID,
				OwnerID:          uintPtr(ownerID),
				CountryID:        countryID,
				FatherNameTemp:   cleanParentName(cat.FatherName),
				MotherNameTemp:   cleanParentName(cat.MotherName),

			}

			result = dbNew.Create(&catGorm)
			if result.Error != nil {
				log.Printf("Error inserting cat %s: %v\n", cat.Name, result.Error)
				continue
			}
			id := catGorm.ID
			if len(titles) > 0 {
				log.Println("titles: ", titles)
				log.Println("catId: ", id)
				insertTitles(dbNew, id, titles)
			}

			newCats = append(newCats, cat.Name)
		}

		offset += len(catMigrations)
	}

	log.Printf("%d cats migrated successfully.\n", len(newCats))
	log.Printf("%d cats have no title.\n", len(noTitleCats))
}

func UpdateCatParents(db *gorm.DB) error {
    // 1. Consulte todos os gatos na tabela
    var cats []Cat
    result := db.Find(&cats)
    if result.Error != nil {
        return result.Error
    }

    count := 0 // Inicialize o contador

    // 2. Itere sobre todos os gatos e atualize os IDs dos pais na tabela
    for _, cat := range cats {
        // Busque o ID do pai e da mãe na tabela cats usando os nomes limpos
        var fatherID, motherID *uint
		
		// Se o nome do pai/mãe temporário não estiver vazio, tente buscar o registro correspondente
		if cat.FatherNameTemp != "" {
			var father Cat
			result := db.Where("name LIKE ?", "%" + cat.FatherNameTemp + "%").First(&father)
			if result.Error != nil {
				if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
					// Se o erro for diferente de 'não encontrado', retorne o erro
					return result.Error
				}
				// Se não encontrar o pai, deixe fatherID como nil
				fatherID = nil
			} else {
				// Se encontrar, use o ID do pai
				fatherID = &father.ID
			}
		}

// Repita a lógica similar para motherID


		if cat.MotherNameTemp != "" {
			var mother Cat
			result := db.Where("name LIKE ?", "%" + cat.MotherNameTemp + "%").First(&mother)
			if result.Error != nil {
				if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
					// Se ocorrer um erro diferente de "não encontrado", retorne o erro
					return result.Error
				}
				// Se não encontrar a mãe, deixe motherID como nil
				// Isso é crucial para assegurar que motherID possa ser definido como nil
				motherID = nil
			} else {
				// Se encontrar a mãe, atribua o ID dela ao motherID
				motherID = &mother.ID
			}
		}


        // Atualize o registro correspondente na tabela cats com os IDs dos pais
        if fatherID != nil || motherID != nil {
            if fatherID != nil {
                cat.FatherID = fatherID
            }
            if motherID != nil {
                cat.MotherID = motherID
            }
            result = db.Save(&cat)
            if result.Error != nil {
                return result.Error
            }

            count++ // Incremente o contador
        }
    }

    fmt.Printf("Número de gatos atualizados: %d\n", count) // Imprime o contador

    return nil
}









