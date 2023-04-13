package cattery

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatteryMongo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	BreederName string             `bson:"breederName"`
	OwnerID     primitive.ObjectID `bson:"ownerID"`
	CountryID   primitive.ObjectID `bson:"countryId"`
}
