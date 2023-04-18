package country

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountryRepository struct {
	Client *mongo.Client
	Logger *logrus.Logger
}

func NewCountryRepository(client *mongo.Client, logger *logrus.Logger) *CountryRepository {
	return &CountryRepository{
		Client: client,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "countries"

func (r *CountryRepository) GetAllCountries() ([]CountryMongo, error) {
	r.Logger.Infof("Repository GetAllCountries")
	var countries []CountryMongo
	ctx := context.Background()

	filter := bson.M{"isActivated": true}
	cur, err := r.Client.Database(database).Collection(collection).Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &countries); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllCountries OK")
	return countries, nil
}
