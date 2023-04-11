package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/federation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateFederations(db *gorm.DB, client *mongo.Client) error {
	var federations []*federation.Federation
	if err := db.Unscoped().Find(&federations).Error; err != nil {
		return err
	}

	federationMongos := make([]interface{}, len(federations))
	for i, f := range federations {
		countryId, err := findCountryIdByCode(client, f.CountryCode)
		if err != nil {
			return err
		}

		federationMongos[i] = federation.FederationMongo{
			ID:             primitive.NewObjectID(),
			Name:           f.FederationName,
			FederationCode: f.FederationCode,
			CountryId:      countryId,
		}
	}

	_, err := client.Database("amacoon").Collection("federations").InsertMany(context.Background(), federationMongos)
	if err != nil {
		return err
	}

	return nil
}
