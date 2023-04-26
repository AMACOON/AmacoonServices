package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"fmt"
)

func MigrateCattery(db *gorm.DB, client *mongo.Client) error {
	
	fmt.Println("Entrou Migrate Cattery")
	// Busque todos os registros da tabela "gatis" usando GORM
	newCattery := []string{} 
	var breeders []*sql.Cattery
	if err := db.Unscoped().Find(&breeders).Error; err != nil {
		return err
	}

	catteryCollection := client.Database("amacoon").Collection("catteries")

	for _, b := range breeders {
		filter := bson.M{"name": b.BreederName}
		count, err := catteryCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			countryId, err := findCountryIdByCode(client, b.BreederCountry)
			if err != nil {
				return err
			}

			catteryMongo := cattery.CatteryMongo{
				ID:          primitive.NewObjectID(),
				Name:        b.BreederName,
				BreederName: b.BreederOwner,
				OwnerID:     primitive.NilObjectID,
				CountryID:   countryId,
			}

			_, err = catteryCollection.InsertOne(context.Background(), catteryMongo)
			if err != nil {
				return err
			}
			newCattery = append(newCattery, b.BreederName)
		}
	}
	fmt.Println("New cattery: %v\n", newCattery)
	fmt.Println("FIM Migrate Cattery")
	return nil
}
