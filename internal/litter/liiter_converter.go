package litter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LitterRequest) ToLitter() (*Litter, error) {
	motherData, err := lr.MotherData.ToCatService()
	if err != nil {
		return nil, err
	}

	fatherData, err := lr.FatherData.ToCatService()
	if err != nil {
		return nil, err
	}

	motherOwner, err := lr.MotherOwner.ToOwnerService()
	if err != nil {
		return nil, err
	}

	fatherOwner, err := lr.FatherOwner.ToOwnerService()
	if err != nil {
		return nil, err
	}

	requesterID, err := primitive.ObjectIDFromHex(lr.RequesterID)
	if err != nil {
		return nil, err
	}

	kittenData := make([]KittenLitter, len(lr.KittenData))
	for i, kitten := range lr.KittenData {
		kittenData[i] = KittenLitter{
			Name:       kitten.Name,
			Gender:     kitten.Gender,
			BreedName:  kitten.BreedName,
			ColorName:  kitten.ColorName,
			EmsCode:    kitten.EmsCode,
			ColorNameX: kitten.ColorNameX,
			Microchip:  kitten.Microchip,
			Breeding:   kitten.Breeding,
		}
	}

	return &Litter{
		MotherData:     motherData,
		FatherData:     fatherData,
		MotherOwner:    motherOwner,
		FatherOwner:    fatherOwner,
		BirthData:      lr.BirthData.ToBirthLitter(),
		Status:         lr.Status,
		ProtocolNumber: lr.ProtocolNumber,
		RequesterID:    requesterID,
		KittenData:     kittenData,
	}, nil
}

func (blr *BirthLitterRequest) ToBirthLitter() BirthLitter {
	return BirthLitter{
		CatteryName: blr.CatteryName,
		NumKittens:  blr.NumKittens,
		BirthDate:   blr.BirthDate,
		CountryCode: blr.CountryCode,
	}
}
