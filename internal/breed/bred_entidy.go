package breed

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BreedMongo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	BreedCode      string             `bson:"breedCode"`
	BreedName    string             `bson:"breedName"`
	BreedCategory int                `bson:"breedCategory"`
	BreedByGroup string             `bson:"breedByGroup"`
}

type BreedCompatibilityMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	BreedCode1 string             `bson:"breedCode1"`
	BreedCode2 string             `bson:"breedCode2"`
}
