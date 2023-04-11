package country

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CountryMongo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Code         string             `bson:"code"`
	Name         string             `bson:"name"`
	IsActivated  bool               `bson:"isActivated"`
}
