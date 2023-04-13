package migrate

import (
	"context"
	"fmt"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func PopulateCountries(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Entrou Migrate Country")
	var countries []*sql.Country
	if err := db.Unscoped().Find(&countries).Error; err != nil {
		return err
	}

	countryCollection := client.Database("amacoon").Collection("countries")

	for _, c := range countries {
		filter := bson.M{"code": c.CountryCode}
		count, err := countryCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			countryMongo := &country.CountryMongo{
				Code:        c.CountryCode,
				Name:        c.CountryName,
				IsActivated: c.Activate == "s",
			}

			_, err := countryCollection.InsertOne(context.Background(), countryMongo)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("FIM Migrate Country")
	return nil
}

func MigrateBreeds(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Entrou Migrate Breeds")
	var breeds []*sql.Breed
	if err := db.Unscoped().Find(&breeds).Error; err != nil {
		return err
	}

	breedCollection := client.Database("amacoon").Collection("breeds")

	for _, b := range breeds {
		filter := bson.M{"breedCode": b.BreedID}
		count, err := breedCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			breedMongo := breed.BreedMongo{
				ID:            primitive.NewObjectID(),
				BreedCode:     b.BreedID,
				BreedName:     b.BreedName,
				BreedCategory: b.BreedCategory,
				BreedByGroup:  b.BreedByGroup,
			}

			_, err := breedCollection.InsertOne(context.Background(), breedMongo)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("FIM Migrate Brreds")
	return nil
}

func MigrateColors(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Entrou Migrate Colors")
	var colors []*sql.Color
	if err := db.Unscoped().Find(&colors).Error; err != nil {
		return err
	}

	colorCollection := client.Database("amacoon").Collection("colors")

	for _, c := range colors {
		filter := bson.M{"emsCode": c.EmsCode}
		count, err := colorCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			colorMongo := color.ColorMongo{
				ID:        primitive.NewObjectID(),
				BreedCode: c.BreedID,
				EmsCode:   c.EmsCode,
				Name:      c.ColorName,
				Group:     c.Group,
				SubGroup:  c.SubGroup,
			}

			_, err := colorCollection.InsertOne(context.Background(), colorMongo)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("FIM Migrate Colors")
	return nil
}
