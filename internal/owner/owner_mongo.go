package owner

import "go.mongodb.org/mongo-driver/bson/primitive"

type OwnerMongo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	Name         string             `bson:"name"`
	Address      string             `bson:"address"`
	City         string             `bson:"city"`
	State        string             `bson:"state"`
	ZipCode      string             `bson:"zipCode"`
	CountryId    primitive.ObjectID `bson:"countryId"`
	Phone        string             `bson:"phone"`
	Valid        bool               `bson:"valid"`
	Observation  []byte             `bson:"observation"`
	CPF          string             `bson:"cpf"`
}

