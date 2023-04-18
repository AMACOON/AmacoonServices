package migrate

import (
	"context"
	"fmt"

	"github.com/agext/levenshtein"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOwnerIDInCattery(client *mongo.Client, similarityThreshold int) (int, error) {
	ownersCollection := client.Database("amacoon").Collection("owners")
	catteriesCollection := client.Database("amacoon").Collection("catteries")

	// Get all catteries
	catteriesCursor, err := catteriesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	defer catteriesCursor.Close(context.Background())

	var catteries []bson.M
	if err := catteriesCursor.All(context.Background(), &catteries); err != nil {
		return 0, err
	}

	// Get all owners
	ownersCursor, err := ownersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	defer ownersCursor.Close(context.Background())

	var owners []bson.M
	if err := ownersCursor.All(context.Background(), &owners); err != nil {
		return 0, err
	}

	// Compare names and update ownerID if similar
	count := 0
	for _, cattery := range catteries {
		breederName := cattery["breederName"].(string)

		for _, owner := range owners {
			ownerName := owner["name"].(string)

			distance := levenshtein.Distance(ownerName, breederName, nil)
			similarity := 1 - float64(distance)/float64(Max(len(ownerName), len(breederName)))

			if similarity >= float64(similarityThreshold)/100.0 {
				// Update ownerID in catteries
				filter := bson.M{"_id": cattery["_id"]}
				update := bson.M{"$set": bson.M{"ownerID": owner["_id"]}}

				_, err := catteriesCollection.UpdateOne(context.Background(), filter, update)
				if err != nil {
					return count, err
				}
				fmt.Printf("Updated ownerID for cattery %v with owner %v\n", cattery["_id"], owner["_id"])
				count++
				break
			}
		}
	}

	return count, nil
}

// Max retorna o maior valor entre a e b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
