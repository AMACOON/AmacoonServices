package litter

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/service"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Litter struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	MotherData     service.CatService   `bson:"motherData"`
	FatherData     service.CatService   `bson:"fatherData"`
	MotherOwner    service.OwnerService `bson:"motherOwner"`
	FatherOwner    service.OwnerService `bson:"fatherOwner"`
	BirthData      BirthLitter          `bson:"birthData"`
	Status         string               `bson:"status"`
	ProtocolNumber string               `bson:"protocolNumber"`
	RequesterID    primitive.ObjectID   `bson:"requesterID"` // OwnerId q esta logado, Pegar dado na Femea
	KittenData     []KittenLitter       `bson:"kittenData"`
	Files          []utils.Files        `bson:"files"`
}

type BirthLitter struct {
	CatteryName string    `bson:"catteryName"` // Pegar da Femea
	NumKittens  int       `bson:"numKittens"`
	BirthDate   time.Time `bson:"birthDate"`
	CountryCode string    `bson:"countryCode"`
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
