package migrate

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateBreeds(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Iniciando migração de raças")

	var breeds []*sql.Breed
	err := populateCollection(db, client, "breeds", &breeds, func(item interface{}) interface{} {
		b := item.(*sql.Breed)
		return &breed.BreedMongo{
			ID:            primitive.NewObjectID(),
			BreedCode:     b.BreedID,
			BreedName:     b.BreedName,
			BreedCategory: b.BreedCategory,
			BreedByGroup:  b.BreedByGroup,
		}
	}, func(items interface{}) int {
		return len(items.([]*sql.Breed))
	}, func(item interface{}) bson.M {
		b := item.(*sql.Breed)
		return bson.M{"breedCode": b.BreedID}
	})

	if err != nil {
		return err
	}

	fmt.Println("Migração de raças concluída")
	return nil
}
