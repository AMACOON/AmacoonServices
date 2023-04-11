package breeader

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BreederMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"breedCode"`
	OwnerID   primitive.ObjectID `bson:"ownerID"`
	CountryId primitive.ObjectID `bson:"countryId"`
}
