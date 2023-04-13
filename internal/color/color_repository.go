package color

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ColorRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewColorRepository(db *mongo.Client, logger *logrus.Logger) *ColorRepository {
	return &ColorRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "colors"

func (r *ColorRepository) GetAllColorsByBreed(breedCode string) ([]ColorMongo, error) {
	r.Logger.Infof("Repository GetAllColorsByBreed")
	var colors []ColorMongo
	filter := bson.M{"breedCode": breedCode}
	cur, err := r.DB.Database(database).Collection(collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var color ColorMongo
		err := cur.Decode(&color)
		if err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllColorsByBreed OK")
	return colors, nil
}
