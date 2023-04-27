package titlerecognition

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TitleRecognitionRequest struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CatData        CatTitleRequest           `bson:"catData"`
	OwnerData      OwnerTitleRequest         `bson:"ownerData"`
	TitleID        string `bson:"titleID"`
	TitleCode      string             `bson:"titleCode"`
	TitleName      string             `bson:"titleName"`
	Certificate    string             `bson:"certificate"`
	Date           time.Time          `bson:"date"`
	Judge          string             `bson:"judge"`
	Status         string             `bson:"status"`
	ProtocolNumber string             `bson:"protocolNumber"`
	RequesterID    string `bson:"requesterID"`
	Files          []utils.FilesReq      `json:"files"`
}

type CatTitleRequest struct {
	ID           string `bson:"id"`
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

type OwnerTitleRequest struct {
	ID          string `bson:"id"`
	Name        string             `bson:"name"`
	CPF         string             `bson:"cpf"`
	Address     string             `bson:"address"`
	City        string             `bson:"city"`
	State       string             `bson:"state"`
	ZipCode     string             `bson:"zipCode"`
	CountryName string             `bson:"countryName"`
	Phone       string             `bson:"phone"`
}
