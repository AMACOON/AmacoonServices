package litter

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/service"
)

type LitterRequest struct {
	MotherData     service.CatServiceRequest   `json:"motherData"`
	FatherData     service.CatServiceRequest   `json:"fatherData"`
	MotherOwner    service.OwnerServiceRequest `json:"motherOwner"`
	FatherOwner    service.OwnerServiceRequest `json:"fatherOwner"`
	BirthData      BirthLitterRequest          `json:"birthData"`
	Status         string                      `json:"status"`
	ProtocolNumber string                      `json:"protocolNumber"`
	RequesterID    string                      `json:"requesterID"`
	KittenData     []KittenLitterRequest       `json:"kittenData"`
}

type KittenLitterRequest struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	BreedName  string `json:"breedName"`
	ColorName  string `json:"colorName"`
	EmsCode    string `json:"emsCode"`
	ColorNameX string `json:"colorNameX"`
	Microchip  string `json:"microchip"`
	Breeding   bool   `json:"breeding"`
}

type BirthLitterRequest struct {
	CatteryName string    `json:"catteryName"`
	NumKittens  int       `json:"numKittens"`
	BirthDate   time.Time `json:"birthDate"`
	CountryCode string    `json:"countryCode"`
}
