package litter

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Litter struct {
	MotherData     CatLitter           `bson:"motherData"`
	FatherData     CatLitter           `bson:"fatherData"`
	BirthData      BirthLitter         `bson:"birthData"`
	KittenData     []KittenLitter      `bson:"kittenData"`
	Status         string              `bson:"status"`
	ProtocolNumber string              `bson:"protocolNumber"`
	RequesterID    primitive.ObjectID  `bson:"requesterID"` // OwnerId q esta logado, Pegar dado na Femea
	Files          []utils.Files       `bson:"files"`
}

type CatLitter struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Registration string             `bson:"registration"`
	Microchip    string             `bson:"microchip"`
	BreedName    string             `bson:"breedName"`
	EmsCode      string             `bson:"emsCode"`
	ColorName    string             `bson:"colorName"`
	Gender       string             `bson:"gender"`
	Owner        OwnerLitter        `bson:"owner"`
}

type OwnerLitter struct {
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

type BirthLitter struct {
	CatteryName  string    `bson:"catteryName"` // Pegar da Femea
	NumKittens   int       `bson:"numKittens"`
	BirthDate    time.Time `bson:"birthDate"`
	CountryCode  string    `bson:"countryCode"`
}

type KittenLitter struct {
	Name       string `bson:"name"`
	Gender     string `bson:"gender"`
	BreedName  string `bson:"breedName"`
	ColorName  string `bson:"colorName"`
	EmsCode    string `bson:"emsCode"`
	ColorNameX string `bson:"colorNameX"`
	Microchip  string `bson:"microchip"`
	Breeding   bool   `bson:"breeding"`
}
