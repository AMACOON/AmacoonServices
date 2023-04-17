package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
)

type TransferRequest struct {
	CatData        CatTransferRequest      `json:"catData"`
	SellerData     SellerTransferRequest   `json:"sellerData"`
	BuyerData      BuyerTransferRequest    `json:"buyerData"`
	Status         string           `json:"status"`
	ProtocolNumber string           `json:"protocolNumber"`
	RequesterID    string           `json:"requesterID"`
	Files          []utils.FilesReq `json:"files"`
}

type CatTransferRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Registration string `json:"registration"`
	Microchip    string `json:"microchip"`
	BreedName    string `json:"breedName"`
	EmsCode      string `json:"emsCode"`
	ColorName    string `json:"colorName"`
	Gender       string `json:"gender"`
	FatherName   string `json:"fatherName"`
	MotherName   string `json:"motherName"`
}

type SellerTransferRequest struct {
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

type BuyerTransferRequest struct {
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
