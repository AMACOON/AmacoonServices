package migrate

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateOwners(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Iniciando migração de proprietários")

	var owners []*sql.Owner
	err := populateCollection(db, client, "owners", &owners, func(item interface{}) interface{} {
		o := item.(*sql.Owner)
		countryId, err := findCountryIdByCode(client, o.Country)
		if err != nil {
			return err
		}

		return &owner.OwnerMongo{
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
	}, func(items interface{}) int {
		return len(items.([]*sql.Owner))
	}, func(item interface{}) bson.M {
		o := item.(*sql.Owner)
		return bson.M{"email": o.Email}
	})

	if err != nil {
		return err
	}

	fmt.Println("Migração de proprietários concluída")
	return nil
}
