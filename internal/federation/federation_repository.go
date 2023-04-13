package federation

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FederationRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewFederationRepository(db *mongo.Client, logger *logrus.Logger) *FederationRepository {
	return &FederationRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "catteries"

func (r *FederationRepository) GetFederationByID(id string) (*FederationMongo, error) {
	r.Logger.Infof("Repository GetFederationByID")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err

	}
	filter := bson.M{"_id": objID}
	var federation FederationMongo
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&federation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithError(err).Errorf("federation not found")
			return nil, err
		}
		return nil, err
	}
	r.Logger.Infof("Repository GetFederationByID OK")
	return &federation, nil
}

func (r *FederationRepository) GetAllFederations() ([]FederationMongo, error) {
    r.Logger.Infof("Repository GetAllFederations")
	var federations []FederationMongo
    cursor, err := r.DB.Database(database).Collection(collection).Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(context.Background(), &federations); err != nil {
        r.Logger.WithError(err).Errorf("error with cursor: %v", err)
		return nil, err
    }
	r.Logger.Infof("Repository GetAllFederations OK")
    return federations, nil
}
