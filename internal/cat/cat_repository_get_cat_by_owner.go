package cat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CatRepository) GetAllByOwnerAndGender(ownerID, sex string) ([]*CatComplete, error) {
	r.Logger.Infof("Repository GetAllByOwnerAndSex")

	catCollection := r.DB.Database(database).Collection(catsCollection)

	ownerObjectID, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid owner ID: %v", ownerID)
		return nil, err
	}

	matchStage := bson.D{{
		Key: "$match",
		Value: bson.M{
			"ownerId": ownerObjectID,
			"sex":     sex,
		},
	}}

	// Pass the required lookups as a slice of strings
	lookups := []string{"color", "breed", "mother", "cattery", "country", "federation", "owner", "father"}

	pipeline := BuildPipelineWithLookups(matchStage, lookups)

	cursor, err := catCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		r.Logger.WithError(err).Error("error getting cats")
		return nil, err
	}
	defer cursor.Close(context.Background())

	catsComplete := []*CatComplete{}
	for cursor.Next(context.Background()) {
		catComplete := &CatComplete{}
		err := cursor.Decode(catComplete)
		if err != nil {
			r.Logger.WithError(err).Error("error decoding cat")
			return nil, err
		}
		catsComplete = append(catsComplete, catComplete)
	}

	if err := cursor.Err(); err != nil {
		r.Logger.WithError(err).Error("cursor error")
		return nil, err
	}

	return catsComplete, nil
}
