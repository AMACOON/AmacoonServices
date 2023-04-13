package owner

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OwnerRepository struct {
	Client *mongo.Client
	Logger *logrus.Logger
}

func NewOwnerRepository(client *mongo.Client, logger *logrus.Logger) *OwnerRepository {
	return &OwnerRepository{
		Client: client,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "owners"

func (r *OwnerRepository) GetOwnerByExhibitorID(id string) (*OwnerMongo, error) {
	r.Logger.Infof("Repository GetOwnerByExhibitorID")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err

	}
	filter := bson.M{"_id": objID}
	var owner OwnerMongo
	err = r.Client.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&owner)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithError(err).Errorf("Owner not found")
			return nil, err
		}
		return nil, err
	}
	r.Logger.Infof("Repository GetBreedByID OK")
	return &owner, nil
}

func (r *OwnerRepository) GetAllOwners() ([]OwnerMongo, error) {
    r.Logger.Infof("Repository GetAllOwners")
	var owners []OwnerMongo
    cursor, err := r.Client.Database(database).Collection(collection).Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(context.Background(), &owners); err != nil {
        r.Logger.WithError(err).Errorf("error with cursor: %v", err)
		return nil, err
    }
	r.Logger.Infof("Repository GetAllOwners OK")
    return owners, nil
}

