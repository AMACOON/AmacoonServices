package migrate

import (
	"context"
	"fmt"
	"strings"

	"github.com/agext/levenshtein"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongomodels "github.com/scuba13/AmacoonServices/config/migrate/models/mongo"
)





func UpdateCatsWithParentIDs3(mongoClient *mongo.Client) error {
	ctx := context.Background()

	// Connect to the cats_temp collection
	mongoDB := mongoClient.Database("amacoon")
	// Create collection variables for both collections
	catsTempCollection := mongoDB.Collection("cats_temp")
	catsCollection := mongoDB.Collection("cats")

	// Prepare a cursor to iterate over all documents in the cats_temp collection
	cursor, err := catsTempCollection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return fmt.Errorf("failed to find cats_temp: %w", err)
	}
	defer cursor.Close(ctx)

	// Get all cat names
	catCursor, err := catsCollection.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to find cats_temp: %w", err)
	}
	defer catCursor.Close(ctx)

	var allCats []mongomodels.Cat
	if err := catCursor.All(ctx, &allCats); err != nil {
		return fmt.Errorf("failed to get all cats: %w", err)
	}

	// Variables to count fathers and mothers
	totalFathers := 0
	foundFathers := 0
	totalMothers := 0
	foundMothers := 0

	// Iterate over each document in the cats_temp collection
	for cursor.Next(ctx) {
		var catTemp mongomodels.CatTempFull
		err := cursor.Decode(&catTemp)
		if err != nil {
			return fmt.Errorf("failed to decode cat_temp: %w", err)
		}

		if catTemp.FatherName != "" {
			totalFathers++
			fatherName := strings.Join(strings.Fields(strings.TrimSpace(catTemp.FatherName)), " ")
			bestMatch := primitive.NilObjectID
			bestDistance := 3
			for _, cat := range allCats {
				distance := levenshtein.Distance(cat.Name, fatherName, nil)
				if distance < bestDistance {
					bestDistance = distance
					bestMatch = cat.ID
				}
			}
			catTemp.FatherID = bestMatch
			foundFathers++
		}
		if catTemp.MotherName != "" {
			totalMothers++
			motherName := strings.Join(strings.Fields(strings.TrimSpace(catTemp.MotherName)), " ")

			bestMatch := primitive.NilObjectID
			bestDistance := 3
			for _, cat := range allCats {
				distance := levenshtein.Distance(cat.Name, motherName, nil)
				if distance < bestDistance {
					bestDistance = distance
					bestMatch = cat.ID
				}
			}
			catTemp.MotherID = bestMatch
			foundMothers++
		}
		// Update the cat_temp document with FatherID and MotherID
		_, err = catsTempCollection.UpdateOne(ctx, bson.M{"_id": catTemp.ID}, bson.M{"$set": bson.M{"fatherID": catTemp.FatherID, "motherID": catTemp.MotherID}})
		if err != nil {
			return fmt.Errorf("failed to update cat_temp: %w", err)
		}
	}

	fmt.Printf("Total fathers in cats_temp: %d, found fathers: %d\n", totalFathers, foundFathers)
	fmt.Printf("Total mothers in cats_temp: %d, found mothers: %d\n", totalMothers, foundMothers)

	return nil
}
