package migrate

import (
	"context"
	"log"

	model "github.com/scuba13/AmacoonServices/internal/title"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTitles(client *mongo.Client) {
	titles := []model.TitlesMongo{
		// Championship/Premiorship Titles
		{Name: "Champion", Code: "CH", Type: "Championship/Premiorship Titles"},
		{Name: "Premior", Code: "PR", Type: "Championship/Premiorship Titles"},
		{Name: "International Champion", Code: "IC", Type: "Championship/Premiorship Titles"},
		{Name: "International Premior", Code: "IP", Type: "Championship/Premiorship Titles"},
		{Name: "Grand International Champion", Code: "GIC", Type: "Championship/Premiorship Titles"},
		{Name: "Grand International Premior", Code: "GIP", Type: "Championship/Premiorship Titles"},
		{Name: "Supreme Champion", Code: "SC", Type: "Championship/Premiorship Titles"},
		{Name: "Supreme Premior", Code: "SP", Type: "Championship/Premiorship Titles"},
		// Winner Titles
		{Name: "Junior Winner", Code: "JW", Type: "Winner Titles"},
		{Name: "DSW", Code: "DSW", Type: "Winner Titles"},
		{Name: "National Winner", Code: "NW", Type: "Winner Titles"},
		{Name: "AW", Code: "AW", Type: "Winner Titles"},
		{Name: "BW", Code: "BW", Type: "Winner Titles"},
		{Name: "CEW", Code: "CEW", Type: "Winner Titles"},
		{Name: "MW", Code: "MW", Type: "Winner Titles"},
		{Name: "NSW", Code: "NSW", Type: "Winner Titles"},
		{Name: "SW", Code: "SW", Type: "Winner Titles"},
		{Name: "World Winner", Code: "WW", Type: "Winner Titles"},
		// Merit Titles
		{Name: "Distinguished Merit", Code: "DM", Type: "Merit Titles"},
		{Name: "DSM", Code: "DSM", Type: "Merit Titles"},
		{Name: "DVM", Code: "DVM", Type: "Merit Titles"},
	}

	// Convertendo o slice de TitlesMongo para um slice de interface{}
	interfaceSlice := make([]interface{}, len(titles))
	for i, title := range titles {
		interfaceSlice[i] = title
	}

	collection := client.Database("amacoon").Collection("titles")
	_, err := collection.InsertMany(context.Background(), interfaceSlice)
	if err != nil {
		log.Fatal(err)
	}
}
