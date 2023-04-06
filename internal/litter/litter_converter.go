package litter

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type LitterConverter struct {
	Logger *logrus.Logger
}

func NewLitterConverter(logger *logrus.Logger) *LitterConverter {
	return &LitterConverter{
		Logger:          logger,
	}
}

func (c *LitterConverter) LitterDBToLitter(litter *LitterDB, kittens []*KittenDB, filesDB []*utils.FilesDB) *Litter {
	litterData := &Litter{
		MotherData: CatLitter{
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
		FatherData: CatLitter{
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
		BirthData: BirthLitter{
			CatteryID:   litter.CatteryID,
			CatteryName: litter.CatteryName,
			NumKittens:  litter.NumKittens,
			BirthDate:   litter.BirthDate,
			Country:     litter.Country,
		},
		LitterID:       litter.ID,
		Status:         litter.Status,
		ProtocolNumber: litter.ProtocolNumber,
	}

	kittenDataSlice := make([]KittenLitter, len(kittens))
	for i, kitten := range kittens {
		kittenData := KittenLitter{
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

	fileSlice := make([]utils.Files, len(filesDB))
	for i, file := range filesDB {
		fileData := utils.Files{
			ID:             file.ID,
			Name:           file.Name,
			Type:           file.Type,
			Base64:         file.Base64,
			ProtocolNumber: file.ProtocolNumber,
			ServiceID:      file.ServiceID,
		}
		fileSlice[i] = fileData
	}
	litterData.Files = fileSlice

	return litterData
}

func (c *LitterConverter) LitterToLitterDB(litterData Litter) (LitterDB, []*KittenDB, []*utils.FilesDB) {

	litter := LitterDB{
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

		CatteryID:      litterData.BirthData.CatteryID,
		CatteryName:    litterData.BirthData.CatteryName,
		NumKittens:     litterData.BirthData.NumKittens,
		BirthDate:      litterData.BirthData.BirthDate,
		Country:        litterData.BirthData.Country,
		ProtocolNumber: litterData.ProtocolNumber,
		Status:         litterData.Status,
	}

	// Transform each KittenData into a models.Kitten struct
	kittenPointers := make([]*KittenDB, 0)
	for _, kittenData := range litterData.KittenData {
		kitten := KittenDB{
			//LitterID:   *kittenData.LitterID,
			//ID:         *kittenData.KittenID,
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

	// Transform each file into a utils.FilesDB struct
	filePointers := make([]*utils.FilesDB, 0)
	for _, fileData := range litterData.Files {
		files := utils.FilesDB{
			//ID : fileData.ID,
			Name:   fileData.Name,
			Type:   fileData.Type,
			Base64: fileData.Base64,
			//ProtocolNumber  : fileData.ProtocolNumber,
			//ServiceID : fileData.ServiceID,
		}
		filePointers = append(filePointers, &files)
	}

	fmt.Println("Convert Litter Create - OK")
	return litter, kittenPointers, filePointers
}
