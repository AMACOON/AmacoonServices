package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/cattery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateCattery(db *gorm.DB, client *mongo.Client) error {
	// Busque todos os registros da tabela "gatis" usando GORM
	var breeders []*cattery.Cattery
	if err := db.Unscoped().Find(&breeders).Error; err != nil {
		return err
	}

	// Converta os registros do GORM para o formato do MongoDB
	breederMongos := make([]interface{}, len(breeders))
	for i, b := range breeders {
		countryId, err := findCountryIdByCode(client, b.BreederCountry)
		if err != nil {
			return err
		}

		breederMongos[i] = cattery.CatteryMongo{
			ID:   primitive.NewObjectID(),
			Name: b.BreederName,
			BreederName: b.BreederOwner,
			OwnerID:   primitive.NilObjectID,
			CountryID: countryId,
		}
	}

	// Insira os registros convertidos na coleção do MongoDB
	_, err := client.Database("amacoon").Collection("catteries").InsertMany(context.Background(), breederMongos)
	if err != nil {
		return err
	}

	return nil
}
