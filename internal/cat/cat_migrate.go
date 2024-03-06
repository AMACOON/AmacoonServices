package cat

import (
	"log"
	//"github.com/scuba13/AmacoonServices/internal/utils"

	"fmt"

	"gorm.io/gorm"
)

const batchSize = 200

func MigrateCats(dbOld *gorm.DB, dbNew *gorm.DB) {
	var catMigrations []CatTable
	var err error
	var offset int

	newCats := []string{}
	noTitleCats := []string{}

	for {
		catMigrations, err = GetCatsMigrate(dbOld, offset, batchSize) //
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

			breedID, err := getBreedID(dbNew, cat.BreedID)
			log.Printf("Cor Gato")
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

			fatherColorID, err := getColorID(dbNew, cat.FatherIdEmscode, cat.FatherBreed)
			if err != nil {
				log.Printf("Error getting getColorID for cat %d: %v\n", i, err)
				continue
			}

			motherColorID, err := getColorID(dbNew, cat.MotherIdEmscode, cat.MotherBreed)
			if err != nil {
				log.Printf("Error getting getColorID for cat %d: %v\n", i, err)
				continue
			}

			fatherBreedID, err := getBreedID(dbNew, cat.FatherBreed)
			if err != nil {
				log.Printf("Error getting getBreedID for cat %d: %v\n", i, err)
				continue
			}

			motherBreedID, err := getBreedID(dbNew, cat.MotherBreed)
			if err != nil {
				log.Printf("Error getting getBreedID for cat %d: %v\n", i, err)
				continue
			}

			catGorm := Cat{
				Name:                cat.Name,
				Registration:        cat.Registration,
				RegistrationType:    cat.RegType,
				Microchip:           cat.Microchip,
				Gender:              sexString,
				Birthdate:           cat.BirthDate,
				Neutered:            &neutered,
				Validated:           validated,
				Observation:         "",
				Fifecat:             fifecat,
				FederationID:        federationID,
				BreedID:             uintPtr(breedID),
				ColorID:             uintPtr(colorID),
				FatherID:            nil,
				MotherID:            nil,
				CatteryID:           catteryID,
				OwnerID:             uintPtr(ownerID),
				CountryID:           countryID,
				FatherNameTemp:      cleanParentName(cat.FatherName),
				MotherNameTemp:      cleanParentName(cat.MotherName),
				FatherNameManual:    &cat.FatherName,
				FatherBreedIDManual: &fatherBreedID,
				FatherColorIDManual: &fatherColorID,
				MotherNameManual:    &cat.MotherName,
				MotherBreedIDManual: &motherBreedID,
				MotherColorIDManual: &motherColorID,
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
		var fatherUpdated, motherUpdated bool

		// Busque e atualize o ID do pai se necessário
		if cat.FatherNameTemp != "" {
			var father Cat
			result := db.Where("name LIKE ?", "%"+cat.FatherNameTemp+"%").First(&father)
			if result.Error != nil {
				// Se não encontrar o pai, deixe fatherID como nil
				cat.FatherID = nil
			} else {
				// Se encontrar, use o ID do pai e limpa os campos temporários
				cat.FatherID = &father.ID
				cat.FatherNameManual = nil
				cat.FatherBreedIDManual = nil
				cat.FatherColorIDManual = nil
				fatherUpdated = true
			}
		}

		// Busque e atualize o ID da mãe se necessário
		if cat.MotherNameTemp != "" {
			var mother Cat
			result := db.Where("name LIKE ?", "%"+cat.MotherNameTemp+"%").First(&mother)
			if result.Error != nil {
				// Se não encontrar a mãe, deixe motherID como nil
				cat.MotherID = nil
			} else {
				// Se encontrar, use o ID da mãe e limpa os campos temporários
				cat.MotherID = &mother.ID
				cat.MotherNameManual = nil
				cat.MotherBreedIDManual = nil
				cat.MotherColorIDManual = nil
				motherUpdated = true
			}
		}

		// Salva as alterações se houve alguma atualização
		if fatherUpdated || motherUpdated {
			result = db.Save(&cat)
			if result.Error != nil {
				return result.Error
			}
			count++
		}
	}

	fmt.Printf("Número de gatos atualizados: %d\n", count) // Imprime o contador

	return nil
}
