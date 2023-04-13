package cattery

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatteryRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewCatteryRepository(db *mongo.Client, logger *logrus.Logger) *CatteryRepository {
	return &CatteryRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "catteries"

func (r *CatteryRepository) GetCatteryByID(id string) (*CatteryMongo, error) {
	r.Logger.Infof("Repository GetCatteryByID")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err

	}
	filter := bson.M{"_id": objID}
	var cattery CatteryMongo
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&cattery)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithError(err).Errorf("breed not found")
			return nil, err
		}
		return nil, err
	}
	r.Logger.Infof("Repository GetBreedByID OK")
	return &cattery, nil
}

func (r *CatteryRepository) GetAllCatteries() ([]CatteryMongo, error) {
    r.Logger.Infof("Repository GetAllCatteries")
	var federations []CatteryMongo
    cursor, err := r.DB.Database(database).Collection(collection).Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(context.Background(), &federations); err != nil {
        r.Logger.WithError(err).Errorf("error with cursor: %v", err)
		return nil, err
    }
	r.Logger.Infof("Repository GetAllCatteries OK")
    return federations, nil
}
