package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (csr *CatServiceRequest) ToCatService() (CatService, error) {
	catID, err := primitive.ObjectIDFromHex(csr.ID)
	if err != nil {
		return CatService{}, err
	}

	return CatService{
		ID:           catID,
		Name:         csr.Name,
		Registration: csr.Registration,
		Microchip:    csr.Microchip,
		BreedName:    csr.BreedName,
		EmsCode:      csr.EmsCode,
		ColorName:    csr.ColorName,
		Gender:       csr.Gender,
		FatherName:   csr.FatherName,
		MotherName:   csr.MotherName,
	}, nil
}

func (osr *OwnerServiceRequest) ToOwnerService() (OwnerService, error) {
	ownerID, err := primitive.ObjectIDFromHex(osr.ID)
	if err != nil {
		return OwnerService{}, err
	}

	return OwnerService{
		ID:          ownerID,
		Name:        osr.Name,
		CPF:         osr.CPF,
		Address:     osr.Address,
		City:        osr.City,
		State:       osr.State,
		ZipCode:     osr.ZipCode,
		CountryName: osr.CountryName,
		Phone:       osr.Phone,
	}, nil
}
