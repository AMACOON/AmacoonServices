package cat

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatMongo struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	Registration     string             `bson:"registration"`
	RegistrationType string             `bson:"registrationType"`
	Microchip        string             `bson:"microchip"`
	Gender           string             `bson:"gender"`
	Birthdate        time.Time          `bson:"birthdate"`
	Neutered         bool               `bson:"neutered"`
	Validated        bool               `bson:"validated"`
	Observation      string             `bson:"observation"`
	Fifecat          bool               `bson:"fifecat"`
	Titles           []TitlesCatsMongo  `bson:"titles"`
	FederationID     primitive.ObjectID `bson:"federationId"`
	BreedID          primitive.ObjectID `bson:"breedId"`
	ColorID          primitive.ObjectID `bson:"colorId"`
	FatherID         primitive.ObjectID `bson:"fatherId"`
	MotherID         primitive.ObjectID `bson:"motherId"`
	CatteryID        primitive.ObjectID `bson:"catteryId"`
	OwnerID          primitive.ObjectID `bson:"ownerId"`
	CountryID        primitive.ObjectID `bson:"countryId"`
	Files            []utils.Files      `bson:"files"`
}

type TitlesCatsMongo struct {
	TitleID      primitive.ObjectID `bson:"id"`
	Date         time.Time          `bson:"date"`
	FederationID primitive.ObjectID `bson:"federationId"`
}
