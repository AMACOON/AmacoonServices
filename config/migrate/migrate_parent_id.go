package migrate

import (
	"context"
	"fmt"

	mongomodels "github.com/scuba13/AmacoonServices/config/migrate/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateCatsTempWithPattensIDs3(client *mongo.Client) error {
	catsTempCollection := client.Database("amacoon").Collection("cats_temp")
	catsCollection := client.Database("amacoon").Collection("cats")

	// Find all documents in cats_temp collection
	count, err := catsTempCollection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("error counting documents in cats_temp collection: %v", err)
	}
	fmt.Printf("Found %d documents in cats_temp collection\n", count)

	cursor, err := catsTempCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("error finding documents in cats_temp collection: %v", err)
	}

	// Iterate over the cursor and update the corresponding document in cats collection
	for cursor.Next(context.Background()) {
		var tempCat mongomodels.CatTempId
		err = cursor.Decode(&tempCat)
		if err != nil {
			return fmt.Errorf("error decoding document from cursor: %v", err)
		}

		fmt.Printf("Processing registration '%s'\n", tempCat.Registro)

		// Find the corresponding cat in the cats collection and update its FatherID and MotherID fields
		filter := bson.M{"registration": tempCat.Registro}
		update := bson.M{"$set": bson.M{"fatherId": tempCat.FatherID, "motherId": tempCat.MotherID}}
		opts := options.Update().SetUpsert(true)
		_, err := catsCollection.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			return fmt.Errorf("error updating cat document: %v", err)
		}

		fmt.Printf("Updated cat with registration '%s'\n", tempCat.Registro)

		fmt.Printf("Updated cat with registration '%s'\n", tempCat.Registro)

	}
	return nil
}
