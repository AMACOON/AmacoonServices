package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Files struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Type           string             `bson:"type"`
	Base64         string             `bson:"base64"`
	ProtocolNumber string             `bson:"protocolNumber"`
}
