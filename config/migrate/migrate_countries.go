package migrate

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/country"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateCountries(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Iniciando migração de países")

	var countries []*sql.Country
	err := populateCollection(db, client, "countries", &countries, func(item interface{}) interface{} {
		c := item.(*sql.Country)
		return &country.CountryMongo{
			Code:        c.CountryCode,
			Name:        c.CountryName,
			IsActivated: c.Activate == "s",
		}
	}, func(items interface{}) int {
		return len(items.([]*sql.Country))
	}, func(item interface{}) bson.M {
		c := item.(*sql.Country)
		return bson.M{"code": c.CountryCode}
	})

	if err != nil {
		return err
	}

	fmt.Println("Migração de países concluída")
	return nil
}
