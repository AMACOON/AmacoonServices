package cat

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatMongo struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name"`
	Registration           string             `bson:"registration"`
	RegistrationType       string             `bson:"registrationType"`
	Microchip              string             `bson:"microchip"`
	Sex                    string             `bson:"sex"`
	Birthdate              string             `bson:"birthdate"`
	Neutered               bool               `bson:"neutered"`
	Validated              bool               `bson:"validated"`
	Titles                 []string           `bson:"titles"`
	RegistrationFederation primitive.ObjectID `bson:"registrationFederation"`
	BreedID                primitive.ObjectID `bson:"breedID"`
	ColorID                primitive.ObjectID `bson:"colorID"`
	FatherID               primitive.ObjectID `bson:"fatherID"`
	MotherID               primitive.ObjectID `bson:"motherID"`
	BreederID              primitive.ObjectID `bson:"breederID"`
	OwnerID                primitive.ObjectID `bson:"ownerID"`
	CountryId              primitive.ObjectID `bson:"countryId"`
}
