package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Files struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Type           string             `bson:"type"`
	Base64         string             `bson:"base64"`
}

type FilesReq struct {
	Name           string             `json:"name"`
	Type           string             `json:"type"`
	Base64         string             `json:"base64"`
}
