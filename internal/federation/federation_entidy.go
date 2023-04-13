package federation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FederationMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	FederationCode string             `bson:"federationCode"`
	CountryId      primitive.ObjectID `bson:"countryId"`
}
