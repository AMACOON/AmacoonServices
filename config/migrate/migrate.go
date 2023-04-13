package migrate

import (
	"context"
	"fmt"

	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
)

func PopulateCountries(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Entrou Migrate")
	var countries []*sql.Country
	if err := db.Unscoped().Find(&countries).Error; err != nil {
		return err
	}

	countryMongos := make([]interface{}, len(countries))
	for i, c := range countries {
		countryMongos[i] = &country.CountryMongo{
			Code:        c.CountryCode,
			Name:        c.CountryName,
			IsActivated: c.Activate == "s",
		}
	}

	_, err := client.Database("amacoon").Collection("countries").InsertMany(context.Background(), countryMongos)
	if err != nil {
		return err
	}
	fmt.Println("FIM Migrate")
	return nil
}

func MigrateBreeds(db *gorm.DB, client *mongo.Client) error {
	var breeds []*sql.Breed
	if err := db.Unscoped().Find(&breeds).Error; err != nil {
		return err
	}

	breedMongos := make([]interface{}, len(breeds))
	for i, b := range breeds {
		breedMongos[i] = breed.BreedMongo{
			ID:            primitive.NewObjectID(),
			BreedCode:     b.BreedID,
			BreedName:     b.BreedName,
			BreedCategory: b.BreedCategory,
			BreedByGroup:  b.BreedByGroup,
		}
	}

	_, err := client.Database("amacoon").Collection("breeds").InsertMany(context.Background(), breedMongos)
	if err != nil {
		return err
	}

	return nil
}

func MigrateColors(db *gorm.DB, client *mongo.Client) error {
	// Busque todos os registros da tabela "cores" usando GORM
	var colors []*sql.Color
	if err := db.Unscoped().Find(&colors).Error; err != nil {
		return err
	}

	// Converta os registros do GORM para o formato do MongoDB
	colorMongos := make([]interface{}, len(colors))
	for i, c := range colors {
		colorMongos[i] = color.ColorMongo{
			ID:        primitive.NewObjectID(),
			BreedCode:   c.BreedID,
			EmsCode:   c.EmsCode,
			Name: c.ColorName,
			Group:     c.Group,
			SubGroup:  c.SubGroup,
		}
	}

	// Insira os registros convertidos na coleção do MongoDB
	_, err := client.Database("amacoon").Collection("colors").InsertMany(context.Background(), colorMongos)
	if err != nil {
		return err
	}

	return nil
}





