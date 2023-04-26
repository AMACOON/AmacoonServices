package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transfer struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CatData        CatTransfer        `bson:"catData"`
	SellerData     OwnerTransfer      `bson:"sellerData"`
	BuyerData      OwnerTransfer      `bson:"buyerData"`
	Status         string             `bson:"status"`
	ProtocolNumber string             `bson:"protocolNumber"`
	RequesterID    primitive.ObjectID `bson:"requesterID"`
	Files          []utils.Files      `bson:"files"`
}

type CatTransfer struct {
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

type OwnerTransfer struct {
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
