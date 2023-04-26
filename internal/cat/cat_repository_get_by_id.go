package cat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *CatRepository) GetCatCompleteByID(id string) (*CatComplete, error) {
	r.Logger.Infof("Repository GetCatCompleteByID")

	catCollection := r.DB.Database("amacoon").Collection("cats")

	catID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid cat ID: %v", id)
		return nil, err
	}
	r.Logger.Infof("Repository GetCatCompleteByID: ", catID)

	matchStage := bson.D{{
		Key: "$match",
		Value: bson.M{
			"_id": catID,
		},
	}}

	// Pass the required lookups as a slice of strings
	lookups := []string{"color", "breed", "cattery", "country", "owner", "federation", "mother", "father", "titles"}

	pipeline := BuildPipelineWithLookups(matchStage, lookups)

	cursor, err := catCollection.Aggregate(context.Background(), pipeline)
	r.Logger.Infof("Aggregate called")
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

		// Adicionar este bloco de código para verificar o primeiro título e modificar "Titles" conforme necessário
		if len(catComplete.Titles) > 0 {
			if catComplete.Titles[0].Title.ID == primitive.NilObjectID {
				catComplete.Titles = []TitlesCats{}
			}

		}

	} else {
		r.Logger.WithField("id", id).Warn("Cat not found")
		return nil, mongo.ErrNoDocuments
	}
	catComplete.FullName = GetFullName (catComplete)
	r.Logger.Infof("Repository GetCatCompleteByID OK")
	return catComplete, nil
}
