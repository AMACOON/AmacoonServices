package color

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ColorMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ColorID   int                `bson:"colorID"`
	BreedID   string             `bson:"breedID"`
	EmsCode   string             `bson:"emsCode"`
	ColorName string             `bson:"colorName"`
	Group     int                `bson:"group"`
	SubGroup  int                `bson:"subGroup"`
}
