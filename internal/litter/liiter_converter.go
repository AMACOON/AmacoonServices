package litter

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func ConvertLitterRequestToLitter(litterReq LitterRequest) (Litter, error) {
    var err error

    // Convertendo o ID do Request
    //var motherDataID, fatherDataID, requesterID primitive.ObjectID
    motherDataID, err := primitive.ObjectIDFromHex(litterReq.MotherData.ID)
    if err != nil {
        return Litter{}, err
    }
    fatherDataID, err := primitive.ObjectIDFromHex(litterReq.FatherData.ID)
    if err != nil {
        return Litter{}, err
    }
    requesterID, err := primitive.ObjectIDFromHex(litterReq.RequesterID)
    if err != nil {
        return Litter{}, err
    }



     // Convertendo os Owners do Request
	 motherOwnerID, err := primitive.ObjectIDFromHex(litterReq.MotherData.Owner.ID)
	 if err != nil {
		 return Litter{}, err
	 }
	 fatherOwnerID, err := primitive.ObjectIDFromHex(litterReq.FatherData.Owner.ID)
	 if err != nil {
		 return Litter{}, err
	 }


    motherOwner := OwnerLitter{
        ID:          motherOwnerID,
        Name:        litterReq.MotherData.Owner.Name,
        CPF:         litterReq.MotherData.Owner.CPF,
        Address:     litterReq.MotherData.Owner.Address,
        City:        litterReq.MotherData.Owner.City,
        State:       litterReq.MotherData.Owner.State,
        ZipCode:     litterReq.MotherData.Owner.ZipCode,
        CountryName: litterReq.MotherData.Owner.CountryName,
        Phone:       litterReq.MotherData.Owner.Phone,
    }

    fatherOwner := OwnerLitter{
        ID:          fatherOwnerID,
        Name:        litterReq.FatherData.Owner.Name,
        CPF:         litterReq.FatherData.Owner.CPF,
        Address:     litterReq.FatherData.Owner.Address,
        City:        litterReq.FatherData.Owner.City,
        State:       litterReq.FatherData.Owner.State,
        ZipCode:     litterReq.FatherData.Owner.ZipCode,
        CountryName: litterReq.FatherData.Owner.CountryName,
        Phone:       litterReq.FatherData.Owner.Phone,
    }


    // Criando a estrutura Litter
    litter := Litter{
        MotherData: CatLitter{
            ID:           motherDataID,
            Name:         litterReq.MotherData.Name,
            Registration: litterReq.MotherData.Registration,
            Microchip:    litterReq.MotherData.Microchip,
            BreedName:    litterReq.MotherData.BreedName,
            EmsCode:      litterReq.MotherData.EmsCode,
            ColorName:    litterReq.MotherData.ColorName,
            Gender:       litterReq.MotherData.Gender,
            Owner:        motherOwner,
        },
        FatherData: CatLitter{
            ID:           fatherDataID,
            Name:         litterReq.FatherData.Name,
            Registration: litterReq.FatherData.Registration,
            Microchip:    litterReq.FatherData.Microchip,
            BreedName:    litterReq.FatherData.BreedName,
            EmsCode:      litterReq.FatherData.EmsCode,
            ColorName:    litterReq.FatherData.ColorName,
            Gender:       litterReq.FatherData.Gender,
            Owner:        fatherOwner,
        },
        BirthData: BirthLitter{
            CatteryName:  litterReq.BirthData.CatteryName,
            NumKittens:   litterReq.BirthData.NumKittens,
            BirthDate:    litterReq.BirthData.BirthDate,
            CountryCode:  litterReq.BirthData.CountryCode,
        },
        KittenData:     convertKittenData(litterReq.KittenData),
        Status:         litterReq.Status,
        ProtocolNumber: litterReq.ProtocolNumber,
        RequesterID:    requesterID,
        Files:          utils.ConvertFilesReqToFiles(litterReq.Files),

    }
fmt.Println(litter)
    return litter, nil
}

func convertKittenData(kittenDataReq []KittenLitterRequest) []KittenLitter {
	kittenData := make([]KittenLitter, len(kittenDataReq))
	for i, kd := range kittenDataReq {
		kittenData[i] = KittenLitter{
			Name:       kd.Name,
			Gender:     kd.Gender,
			BreedName:  kd.BreedName,
			ColorName:  kd.ColorName,
			EmsCode:    kd.EmsCode,
			ColorNameX: kd.ColorNameX,
			Microchip:  kd.Microchip,
			Breeding:   kd.Breeding,
		}
	}
	return kittenData
}


func convertFilesReqToFiles(filesReqList []utils.FilesReq) []utils.Files {
	filesList := make([]utils.Files, len(filesReqList))

	for i, filesReq := range filesReqList {
		filesList[i] = utils.Files{
			ID:       primitive.NewObjectID(), // Gere um novo ObjectID
			Name:     filesReq.Name,
			Type:     filesReq.Type,
			Base64:   filesReq.Base64,
		}
	}

	return filesList
}



