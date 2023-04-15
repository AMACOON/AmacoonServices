package litter

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/utils"
)

type LitterRequest struct {
	MotherData     CatLitterRequest      `json:"motherData"`
	FatherData     CatLitterRequest      `json:"fatherData"`
	BirthData      BirthLitterRequest    `json:"birthData"`
	Status         string                `json:"status"`
	ProtocolNumber string                `json:"protocolNumber"`
	RequesterID    string                `json:"requesterID"` // OwnerId q esta logado, Pegar dado na Femea
	KittenData     []KittenLitterRequest `json:"kittenData"`
	Files          []utils.FilesReq      `json:"files"`
}

type CatLitterRequest struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Registration string         `json:"registration"`
	Microchip    string         `json:"microchip"`
	BreedName    string         `json:"breedName"`
	EmsCode      string         `json:"emsCode"`
	ColorName    string         `json:"colorName"`
	Gender       string         `json:"gender"`
	Owner        OwnerLitterReq `json:"owner"`
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

type OwnerLitterReq struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zipCode"`
	CountryName string `json:"countryName"`
	Phone       string `json:"phone"`
}
