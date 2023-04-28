package title

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/sirupsen/logrus"
)

type TitleRepository struct {
	Client *mongo.Client
	Logger *logrus.Logger
}

func NewTitleRepository(client *mongo.Client, logger *logrus.Logger) *TitleRepository {
	return &TitleRepository{
		Client: client,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "titles"

func (r *TitleRepository) GetAllTitles() ([]TitleMongo, error) {

	r.Logger.Infof("Repository GetAllTitles")
	var titles []TitleMongo
	cursor, err := r.Client.Database(database).Collection(collection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &titles); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllTitles OK")
	return titles, nil
}

