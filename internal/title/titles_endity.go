package title

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TitlesMongo struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Code string             `bson:"code"`
	Type string             `bson:"type"`
	Certificate string      `bson:"certificate"`
	Amount int              `bson:"amount"`
	Observation string     `bson:"observation"`
}