package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateFederations(db *gorm.DB, client *mongo.Client) error {
	var federations []*sql.Federation
	if err := db.Unscoped().Find(&federations).Error; err != nil {
		return err
	}

	federationCollection := client.Database("amacoon").Collection("federations")

	for _, f := range federations {
		filter := bson.M{"federationCode": f.FederationCode}
		count, err := federationCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			countryId, err := findCountryIdByCode(client, f.CountryCode)
			if err != nil {
				return err
			}

			federationMongo := federation.FederationMongo{
				ID:             primitive.NewObjectID(),
				Name:           f.FederationName,
				FederationCode: f.FederationCode,
				CountryId:      countryId,
			}

			_, err = federationCollection.InsertOne(context.Background(), federationMongo)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
