package cat

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type CatAI struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CatID     uint               `bson:"cat_id"`
	Summary   string             `bson:"summary"`
	Embedding []float32          `bson:"embedding"`
}

func generateCatSummary(cat Cat) string {
	fmt.Println("catSummaryName :", cat.Name)
	fmt.Println("catSummaryBreed :", cat.Breed.BreedName)
	fmt.Println("catSummaryCattery :", cat.CatteryID)

	summary := fmt.Sprintf("O gato chamado %s é da raça %s e tem a cor %s. ", cat.Name, cat.Breed.BreedName, cat.Color.Name)
	summary += fmt.Sprintf("O gênero é %s e nasceu em %s. ", cat.Gender, cat.Birthdate.Format("02/01/2006"))

	if cat.FatherName != "" && cat.MotherName != "" {
		summary += fmt.Sprintf("Seu pai é o gato chamado %s e sua mãe é a gata chamada %s. ", cat.FatherName, cat.MotherName)
	}

	if cat.Cattery != nil {
		if cat.Cattery.BreederName != "" && cat.Cattery.Name != "" {
			summary += fmt.Sprintf("O criador é %s e o gatil é chamado de %s. ", cat.Cattery.BreederName, cat.Cattery.Name)
		}
	}

	if cat.Owner != nil && cat.Country != nil {
		summary += fmt.Sprintf("O proprietário do gato é %s e o país de origem é %s. ", cat.Owner.Name, cat.Country.Name)
	}

	if cat.Federation != nil {
		summary += fmt.Sprintf("A federação a qual pertence é a %s. ", cat.Federation.Name)
	}

	if cat.Registration != "" {
		summary += fmt.Sprintf("O numero de registro é %s. ", cat.Registration)
	}

	if cat.Microchip != "" {
		summary += fmt.Sprintf("O microchip implantado é %s.", cat.Microchip)
	}

	return summary
}

func getEmbedding(summary string) []float32 {
	key := "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"
	client := openai.NewClient(key)

	// Configure os parâmetros da chamada de API
	params := openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{summary},
		User:  "Scuba13",
	}

	// Faça a chamada da API para obter os embeddings
	embeddings, err := client.CreateEmbeddings(context.Background(), params)
	if err != nil {
		fmt.Printf("Error fetching embeddings: %v\n", err)
		return nil
	}

	return embeddings.Data[0].Embedding
}

func insertCatAI(mongo *mongo.Client, cat CatAI) error {

	collection := mongo.Database("amacoon").Collection("cats_ai")
	_, err := collection.InsertOne(context.Background(), cat)
	if err != nil {
		return fmt.Errorf("failed to insert CatAI: %v", err)
	}

	return nil
}

func PopulateCatAISummary(db *gorm.DB, mongo *mongo.Client) {
	fmt.Println("Populating cat_ai table...")
	var cats []Cat
	db.Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner").
		Preload("Federation").
		Limit(100).Find(&cats)

	for _, cat := range cats {
		fmt.Println("cat :", cat.Name)
		if cat.FatherID != nil {
			var father Cat
			db.Select("name").Where("id = ?", cat.FatherID).First(&father)
			cat.FatherName = father.Name
		}

		if cat.MotherID != nil {
			var mother Cat
			db.Select("name").Where("id = ?", cat.MotherID).First(&mother)
			cat.MotherName = mother.Name
		}
		summary := generateCatSummary(cat)
		embedding := getEmbedding(summary)

		catAI := CatAI{
			CatID:     cat.ID,
			Summary:   summary,
			Embedding: embedding,
		}

		if err := insertCatAI(mongo, catAI); err != nil {
			fmt.Printf("Error inserting cat_ai record for cat ID %d: %v\n", cat.ID, err)
		} else {
			fmt.Printf("Successfully inserted cat_ai record for cat ID %d\n", cat.ID)
		}
	}
}

func PopulateCatAI(db *gorm.DB, mongo *mongo.Client) {
	fmt.Println("Populating cat_ai table...")
	var cats []Cat
	db.Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner.Country").
		Preload("Federation").
		Preload("Titles.Titles").
		Limit(100).Find(&cats)

	for _, cat := range cats {
		fmt.Println("cat :", cat.Name)
		if cat.FatherID != nil {
			var father Cat
			db.Select("name").Where("id = ?", cat.FatherID).First(&father)
			cat.FatherName = father.Name
		}

		if cat.MotherID != nil {
			var mother Cat
			db.Select("name").Where("id = ?", cat.MotherID).First(&mother)
			cat.MotherName = mother.Name
		}

		catJSON, err := json.Marshal(cat)
		if err != nil {
			fmt.Printf("Error marshalling cat object to JSON for cat ID %d: %v\n", cat.ID, err)
			continue
		}

		embedding := getEmbedding(string(catJSON))

		catAI := CatAI{
			CatID:     cat.ID,
			Summary:   string(catJSON),
			Embedding: embedding,
		}

		if err := insertCatAI(mongo, catAI); err != nil {
			fmt.Printf("Error inserting cat_ai record for cat ID %d: %v\n", cat.ID, err)
		} else {
			fmt.Printf("Successfully inserted cat_ai record for cat ID %d\n", cat.ID)
		}
	}
}
