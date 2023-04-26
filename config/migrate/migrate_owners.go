package migrate

import (
	"context"
	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"fmt"
)

func MigrateOwners(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Entrou Migrate Owners")
	
	var owners []*sql.Owner
	newOwners := []string{}
	ownerCollection := client.Database("amacoon").Collection("owners")
	batchSize := 500
	offset := 0

	for {
		if err := db.Unscoped().Limit(batchSize).Offset(offset).Find(&owners).Error; err != nil {
			return err
		}

		if len(owners) == 0 {
			break
		}

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
				newOwners = append(newOwners, o.OwnerName)
			}
		}

		if len(owners) < batchSize {
			break
		}

		offset += batchSize
	}
	fmt.Println("New owners: %v\n", newOwners)
	fmt.Println("Fim Migrate Owners")
	return nil
}
