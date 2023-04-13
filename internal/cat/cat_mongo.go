package cat

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CatMongo struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Name                     string             `bson:"name"`
	Registration             string             `bson:"registration"`
	RegistrationType         string             `bson:"registrationType"`
	Microchip                string             `bson:"microchip"`
	Sex                      string             `bson:"sex"`
	Birthdate                time.Time             `bson:"birthdate"`
	Neutered                 bool               `bson:"neutered"`
	Validated                bool               `bson:"validated"`
	Observation              string             `bson:"observation"`
	Fifecat                  bool               `bson:"fifecat"`
	Titles                   []string           `bson:"titles"`
	RegistrationFederationID primitive.ObjectID `bson:"registrationFederation"`
	BreedID                  primitive.ObjectID `bson:"breedId"`
	ColorID                  primitive.ObjectID `bson:"colorId"`
	FatherID                 primitive.ObjectID `bson:"fatherId"`
	MotherID                 primitive.ObjectID `bson:"motherId"`
	CatteryID                primitive.ObjectID `bson:"catteryId"`
	OwnerID                  primitive.ObjectID `bson:"ownerId"`
	CountryID                primitive.ObjectID `bson:"countryId"`
}


