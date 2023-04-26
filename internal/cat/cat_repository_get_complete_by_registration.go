package cat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CatRepository) GetCatCompleteByRegistration(registration string) (*CatComplete, error) {
	r.Logger.Infof("Repository GetCatCompleteByRegistration")

	catCollection := r.DB.Database(database).Collection(catsCollection)

	matchStage := bson.D{{
		Key: "$match",
		Value: bson.M{
			"registration": registration,
		},
	}}

	// Pass the required lookups as a slice of strings
	lookups := []string{"color", "breed", "mother", "cattery", "country", "federation", "owner", "father", "titles"}

	pipeline := BuildPipelineWithLookups(matchStage, lookups)

	cursor, err := catCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		r.Logger.WithError(err).Error("error getting cat")
		return nil, err
	}
	defer cursor.Close(context.Background())

	catComplete := &CatComplete{}
	if cursor.Next(context.Background()) {
		err := cursor.Decode(catComplete)
		if err != nil {
			r.Logger.WithError(err).Error("error decoding cat")
			return nil, err
		}
	} else {
		r.Logger.WithField("Registration", registration).Warn("Cat not found")
		return nil, mongo.ErrNoDocuments
	}
	if len(catComplete.Titles) > 0 {
		if catComplete.Titles[0].Title.ID == primitive.NilObjectID {
			catComplete.Titles = []TitlesCats{}
		}
	}
	catComplete.FullName = GetFullName(catComplete)
	r.Logger.Infof("Repository GetCatCompleteByRegistration OK")
	return catComplete, nil
}
