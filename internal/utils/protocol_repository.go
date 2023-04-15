package utils

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProtocolRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewProtocolRepository(db *mongo.Client, logger *logrus.Logger) *ProtocolRepository {
	return &ProtocolRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "protocols"

func (r *ProtocolRepository) ProtocolNumberExists(protocol string) (bool, error) {
	r.Logger.Infof("Repository ProtocolNumberExists")
	filter := bson.M{"protocol": protocol}
	count, err := r.DB.Database(database).Collection(collection).CountDocuments(context.Background(), filter)

	if err != nil {
		return false, err
	}
	r.Logger.Infof("Repository ProtocolNumberExists OK")
	return count > 0, nil
}



func (r *ProtocolRepository) SaveProtocolNumber(protocolNumber string) error {
	r.Logger.Infof("Repository SaveProtocolNumber")

	protocol := Protocol{
		Protocol: protocolNumber,
	}

	_, err := r.DB.Database(database).Collection(collection).InsertOne(context.Background(), protocol)
	if err != nil {
		return err
	}

	r.Logger.Infof("Repository SaveProtocolNumber OK")
	return nil
}


