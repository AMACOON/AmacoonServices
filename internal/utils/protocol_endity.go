package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type Protocol struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Protocol string             `bson:"protocol"`
}
