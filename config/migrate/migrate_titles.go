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
		{Name: "Champion", Code: "CH", Type: "Championship/Premiorship Titles", Certificate: "CAC", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "Premior", Code: "PR", Type: "Championship/Premiorship Titles", Certificate: "CAP", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "International Champion", Code: "IC", Type: "Championship/Premiorship Titles", Certificate: "CACIB", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "International Premior", Code: "IP", Type: "Championship/Premiorship Titles", Certificate: "CAPIB", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "Grand International Champion", Code: "GIC", Type: "Championship/Premiorship Titles", Certificate: "CAGCIB", Amount: 6, Observation: "4 juízes diferentes"},
		{Name: "Grand International Premior", Code: "GIP", Type: "Championship/Premiorship Titles", Certificate: "CAGPIB", Amount: 6, Observation: "4 juízes diferentes"},
		{Name: "Supreme Champion", Code: "SC", Type: "Championship/Premiorship Titles", Certificate: "CACS", Amount: 9, Observation: "5 juízes diferentes"},
		{Name: "Supreme Premior", Code: "SP", Type: "Championship/Premiorship Titles", Certificate: "CAPS", Amount: 9, Observation: "5 juízes diferentes"},
		// Winner Titles
		{Name: "Junior Winner", Code: "JW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "DSW", Code: "DSW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "National Winner", Code: "NW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "AW", Code: "AW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "BW", Code: "BW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "CEW", Code: "CEW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "MW", Code: "MW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "NSW", Code: "NSW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "SW", Code: "SW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "World Winner", Code: "WW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
	
		// Merit Titles
		{Name: "Distinguished Merit", Code: "DM", Type: "Merit Titles", Certificate: "IC/IP para DM", Amount: 5, Observation: "5 certificados de filhos com o título IC/IP ou superior."},
		{Name: "DSM", Code: "DSM", Type: "Merit Titles", Certificate: "BIS para DSM", Amount: 10, Observation: "Em no mínimo 2 anos"},
		{Name: "DVM", Code: "DVM", Type: "Merit Titles", Certificate: "BIV para DVM", Amount: 10, Observation: "Em no mínimo 2 anos"},
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


