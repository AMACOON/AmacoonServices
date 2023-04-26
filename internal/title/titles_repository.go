// This file is a repository for the title entity.
// Make a function that returns a list of all titles by type.
// Path: internal/title/titles_repository.go.
 
package title

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



func GetAllTitlesByType(client *mongo.Client) ([]TitlesMongo, error) {
	var titles []TitlesMongo
	cursor, err := client.Database("amacoon").Collection("titles").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &titles); err != nil {
		return nil, err
	}
	return titles, nil
}

