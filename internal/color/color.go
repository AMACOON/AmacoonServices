package color

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ColorMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BreedCode string             `bson:"breedCode"`
	EmsCode   string             `bson:"emsCode"`
	Name      string             `bson:"name"`
	Group     int                `bson:"group"`
	SubGroup  int                `bson:"subGroup"`
}
