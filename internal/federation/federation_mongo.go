package federation


import "go.mongodb.org/mongo-driver/bson/primitive"

type FederationMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FederationCode string             `bson:"federationCode"`
	Description    string             `bson:"description"`
	CountryId primitive.ObjectID `bson:"countryId"`
}

