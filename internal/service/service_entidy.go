package service

import "go.mongodb.org/mongo-driver/bson/primitive"


type CatService struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Registration string             `bson:"registration"`
	Microchip    string             `bson:"microchip"`
	BreedName    string             `bson:"breedName"`
	EmsCode      string             `bson:"emsCode"`
	ColorName    string             `bson:"colorName"`
	Gender       string             `bson:"gender"`
	FatherName   string             `bson:"fatherName"`
	MotherName   string             `bson:"motherName"`
}

type OwnerService struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	CPF         string             `bson:"cpf"`
	Address     string             `bson:"address"`
	City        string             `bson:"city"`
	State       string             `bson:"state"`
	ZipCode     string             `bson:"zipCode"`
	CountryName string             `bson:"countryName"`
	Phone       string             `bson:"phone"`
}
