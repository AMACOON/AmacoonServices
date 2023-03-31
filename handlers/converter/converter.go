package converter

import (
	"amacoonservices/models"
	"fmt"
)

func TransformLitterAndKittensToLitterData(litter *models.Litter, kittens []*models.Kitten) *models.LitterData {
	litterData := &models.LitterData{
		MotherData: models.CatData{
			Name:         litter.MotherName,
			Registration: litter.MotherReg,
			Microchip:    litter.MotherMicro,
			BreedName:    litter.MotherBreed,
			EmsCode:      litter.MotherEMSCode,
			ColorName:    litter.MotherColor,
			OwnerID:      litter.MotherOwnerID,
			OwnerName:    litter.MotherOwnerName,
			Address:      litter.MotherAddress,
			ZipCode:      litter.MotherZipCode,
			City:         litter.MotherCity,
			State:        litter.MotherState,
			Country:      litter.MotherCountry,
			Phone:        litter.MotherPhone,
		},
		FatherData: models.CatData{
			Name:         litter.FatherName,
			Registration: litter.FatherReg,
			Microchip:    litter.FatherMicro,
			BreedName:    litter.FatherBreed,
			EmsCode:      litter.FatherEMSCode,
			ColorName:    litter.FatherColor,
			OwnerID:      litter.FatherOwnerID,
			OwnerName:    litter.FatherOwnerName,
			Address:      litter.FatherAddress,
			ZipCode:      litter.FatherZipCode,
			City:         litter.FatherCity,
			State:        litter.FatherState,
			Country:      litter.FatherCountry,
			Phone:        litter.FatherPhone,
		},
		BirthData: models.BirthData{
			CatteryID:   litter.CatteryID,
			CatteryName: litter.CatteryName,
			NumKittens:  litter.NumKittens,
			BirthDate:   litter.BirthDate,
			Country:     litter.Country,
		},
		LitterID: litter.ID,
		Status:   litter.Status,
	}

	kittenDataSlice := make([]models.KittenService, len(kittens))
	for i, kitten := range kittens {
		kittenData := models.KittenService{
			KittenID:   &kitten.ID,
			LitterID:   &kitten.LitterID,
			BreedName:  kitten.BreedName,
			ColorName:  kitten.ColorName,
			EmsCodeID:  kitten.EmsCodeID,
			Microchip:  kitten.Microchip,
			ColorNameX: kitten.ColorNameX,
			Breeding:   kitten.Breeding,
			Name:       kitten.Name,
			Sex:        kitten.Sex,
			Status:     kitten.Status,
		}
		kittenDataSlice[i] = kittenData
	}
	litterData.KittenData = kittenDataSlice

	return litterData
}

func TransformLitterDataToLitterAndKittens(litterData models.LitterData) (models.Litter, []*models.Kitten) {
	fmt.Println("Convert Litter Create")
	// Transform LitterData into a models.Litter struct
	litter := models.Litter{
		//LitterID: litterData.LitterID,
		FatherName:      litterData.FatherData.Name,
		FatherReg:       litterData.FatherData.Registration,
		FatherMicro:     litterData.FatherData.Microchip,
		FatherBreed:     litterData.FatherData.BreedName,
		FatherEMSCode:   litterData.FatherData.EmsCode,
		FatherColor:     litterData.FatherData.ColorName,
		FatherOwnerID:   litterData.FatherData.OwnerID,
		FatherOwnerName: litterData.FatherData.OwnerName,
		FatherAddress:   litterData.FatherData.Address,
		FatherZipCode:   litterData.FatherData.ZipCode,
		FatherCity:      litterData.FatherData.City,
		FatherState:     litterData.FatherData.State,
		FatherCountry:   litterData.FatherData.Country,
		FatherPhone:     litterData.FatherData.Phone,

		MotherName:      litterData.MotherData.Name,
		MotherReg:       litterData.MotherData.Registration,
		MotherMicro:     litterData.MotherData.Microchip,
		MotherBreed:     litterData.MotherData.BreedName,
		MotherEMSCode:   litterData.MotherData.EmsCode,
		MotherColor:     litterData.MotherData.ColorName,
		MotherOwnerID:   litterData.MotherData.OwnerID,
		MotherOwnerName: litterData.MotherData.OwnerName,
		MotherAddress:   litterData.MotherData.Address,
		MotherZipCode:   litterData.MotherData.ZipCode,
		MotherCity:      litterData.MotherData.City,
		MotherState:     litterData.MotherData.State,
		MotherCountry:   litterData.MotherData.Country,
		MotherPhone:     litterData.MotherData.Phone,

		CatteryID:   litterData.BirthData.CatteryID,
		CatteryName: litterData.BirthData.CatteryName,
		NumKittens:  litterData.BirthData.NumKittens,
		BirthDate:   litterData.BirthData.BirthDate,
		Country:     litterData.BirthData.Country,

		Status: "Pending",
	}

	// Transform each KittenData into a models.Kitten struct
	kittenPointers := make([]*models.Kitten, 0)
	for _, kittenData := range litterData.KittenData {
		kitten := models.Kitten{
			//LitterID:     litterID,
			//KittenID:    0,
			BreedName:  kittenData.BreedName,
			ColorName:  kittenData.ColorName,
			EmsCodeID:  kittenData.EmsCodeID,
			Microchip:  kittenData.Microchip,
			ColorNameX: kittenData.ColorNameX,
			Breeding:   kittenData.Breeding,
			Name:       kittenData.Name,
			Sex:        kittenData.Sex,
			Status:     kittenData.Status,
		}
		kittenPointers = append(kittenPointers, &kitten)
	}
	fmt.Println("Convert Litter Create - OK")
	return litter, kittenPointers
}
