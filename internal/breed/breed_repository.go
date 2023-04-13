package breed

import (

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "fmt"
    "github.com/sirupsen/logrus"
    "context"
)

type BreedRepository struct {
	DB *mongo.Client
    Logger        *logrus.Logger
}

func NewBreedRepository(db *mongo.Client  ,logger *logrus.Logger) *BreedRepository {
    return &BreedRepository{
        DB: db,
        Logger:       logger,
    }
}

var database= "amacoon"
var collection ="breeds"

func (r *BreedRepository) GetBreedByID(id string) (*BreedMongo, error) {
	r.Logger.Infof("Repository GetBreedByID")
    objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
        r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err

        
	}
	filter := bson.M{"_id": objID}
	var breed BreedMongo
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&breed)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("breed not found")
		}
		return nil, err
	}
	r.Logger.Infof("Repository GetBreedByID OK")
    return &breed, nil
}


func (r *BreedRepository) GetAllBreeds() ([]BreedMongo, error) {
	r.Logger.Infof("Repository GetAllBreeds")
    var breeds []BreedMongo

	collection := r.DB.Database(database).Collection(collection)

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		r.Logger.WithError(err).Error("failed to get all breeds")
		return nil,err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var breed BreedMongo
		err := cursor.Decode(&breed)
		if err != nil {
			r.Logger.WithError(err).Error("error decoding breed: %v", err)
			return nil, err
		}
		breeds = append(breeds, breed)
	}

	if err := cursor.Err(); err != nil {
		r.Logger.WithError(err).Error("error with cursor: %v", err)
		return nil, err
	}
    r.Logger.Infof("Repository GetAllBreeds OK")
	return breeds, nil
}


