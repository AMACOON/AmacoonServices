package converter

import(
	"amacoonservices/models"

)


func TransformLitterAndKittensToLitterData(litter *models.Litter, kittens []*models.Kitten) *models.LitterData {
	litterData := &models.LitterData{
		MotherData: models.CatData{
			Name:          litter.MotherName,
			Registration:  litter.MotherReg,
			Microchip:     litter.MotherMicro,
			BreedName:     litter.MotherBreed,
			EmsCode:     litter.MotherEMSCode,
			ColorName:     litter.MotherColor,
			OwnerID:       litter.MotherOwnerID,
		},
		FatherData: models.CatData{
			Name:          litter.FatherName,
			Registration:  litter.FatherReg,
			Microchip:     litter.FatherMicro,
			BreedName:     litter.FatherBreed,
			EmsCode:     litter.FatherEMSCode,
			ColorName:     litter.FatherColor,
			OwnerID:       litter.FatherOwnerID,
		},
		BirthData: models.BirthData{
			CatteryID:     litter.CatteryID,
			CatteryName:   litter.CatteryName,
			NumKittens:    litter.NumKittens,
			BirthDate:     litter.BirthDate,
		},
		LitterID:   litter.LitterID,
		Status:     litter.Status,
	}

	kittenDataSlice := make([]models.Kitten, len(kittens))
	for i, kitten := range kittens {
		kittenData := models.Kitten{
			KittenID:    kitten.KittenID,
			BreedName:   kitten.BreedName,
			ColorName:   kitten.ColorName,
			EmsCodeID:   kitten.EmsCodeID,
			CountryCode: kitten.CountryCode,
			Microchip:   kitten.Microchip,
			ColorNameX:  kitten.ColorNameX,
			Breeding:    kitten.Breeding,
			Name:        kitten.Name,
			Sex:         kitten.Sex,
			Status:      kitten.Status,
		}
		kittenDataSlice[i] = kittenData
	}
	litterData.KittenData = kittenDataSlice

	return litterData
}

func TransformLitterDataToLitterAndKittens(litterData models.LitterData) (models.Litter, []*models.Kitten) {
	// Transform LitterData into a models.Litter struct
	litter := models.Litter{
		//LitterID:      litterData.LitterID,
		FatherName:    litterData.FatherData.Name,
		FatherReg:     litterData.FatherData.Registration,
		FatherMicro:   litterData.FatherData.Microchip,
		FatherBreed:   litterData.FatherData.BreedName,
		FatherEMSCode: litterData.FatherData.EmsCode,
		FatherColor:   litterData.FatherData.ColorName,
		FatherOwnerID: litterData.FatherData.OwnerID,
		MotherName:    litterData.MotherData.Name,
		MotherReg:     litterData.MotherData.Registration,
		MotherMicro:   litterData.MotherData.Microchip,
		MotherBreed:   litterData.MotherData.BreedName,
		MotherEMSCode: litterData.MotherData.EmsCode,
		MotherColor:   litterData.MotherData.ColorName,
		MotherOwnerID: litterData.MotherData.OwnerID,
		CatteryID:     litterData.BirthData.CatteryID,
		CatteryName:   litterData.BirthData.CatteryName,
		NumKittens:    litterData.BirthData.NumKittens,
		BirthDate:     litterData.BirthData.BirthDate,
		Status:        "Pending",
	}

	// Transform each KittenData into a models.Kitten struct
	kittenPointers := make([]*models.Kitten, 0)
	for _, kittenData := range litterData.KittenData {
		kitten := models.Kitten{
			//LitterID:     litterID,
			//KittenID:    0,
			BreedName:   kittenData.BreedName,
			ColorName:   kittenData.ColorName,
			EmsCodeID:   kittenData.EmsCodeID,
			CountryCode: kittenData.CountryCode,
			Microchip:   kittenData.Microchip,
			ColorNameX:  kittenData.ColorNameX,
			Breeding:    kittenData.Breeding,
			Name:        kittenData.Name,
			Sex:         kittenData.Sex,
			Status:      "Pending",
		}
		kittenPointers = append(kittenPointers, &kitten)
	}

	return litter, kittenPointers
}

