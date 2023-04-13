package owner

import "go.mongodb.org/mongo-driver/bson/primitive"

type OwnerMongo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	Name         string             `bson:"name"`
	CPF          string             `bson:"cpf"`
	Address      string             `bson:"address"`
	City         string             `bson:"city"`
	State        string             `bson:"state"`
	ZipCode      string             `bson:"zipCode"`
	CountryID    primitive.ObjectID `bson:"countryId"`
	Phone        string             `bson:"phone"`
	Valid        bool               `bson:"valid"`
	ValidId      string             `bson:"validId"`
	Observation  string             `bson:"observation"`
}
