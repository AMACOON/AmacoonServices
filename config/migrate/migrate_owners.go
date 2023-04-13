package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateOwners(db *gorm.DB, client *mongo.Client) error {
	var owners []*sql.Owner
	if err := db.Unscoped().Find(&owners).Error; err != nil {
		return err
	}

	ownerCollection := client.Database("amacoon").Collection("owners")

	for _, o := range owners {
		filter := bson.M{"email": o.Email}
		count, err := ownerCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			return err
		}

		if count == 0 {
			countryId, err := findCountryIdByCode(client, o.Country)
			if err != nil {
				return err
			}

			ownerMongo := owner.OwnerMongo{
				ID:           primitive.NewObjectID(),
				Email:        o.Email,
				PasswordHash: o.PasswordHash,
				Name:         o.OwnerName,
				CPF:          o.CPF,
				Address:      o.Address,
				City:         o.City,
				State:        o.State,
				ZipCode:      o.ZipCode,
				CountryID:    countryId,
				Phone:        o.Phone,
				Valid:        o.Valid == "s",
				ValidId:      o.ValidationID,
				Observation:  "",
			}

			_, err = ownerCollection.InsertOne(context.Background(), ownerMongo)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
