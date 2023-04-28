package migrate

import (
	"fmt"

	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"github.com/scuba13/AmacoonServices/internal/color"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateColors(db *gorm.DB, client *mongo.Client) error {
	fmt.Println("Iniciando migração de cores")

	var colors []*sql.Color
	err := populateCollection(db, client, "colors", &colors, func(item interface{}) interface{} {
		c := item.(*sql.Color)
		return &color.ColorMongo{
			ID:        primitive.NewObjectID(),
			BreedCode: c.BreedID,
			EmsCode:   c.EmsCode,
			Name:      c.ColorName,
			Group:     c.Group,
			SubGroup:  c.SubGroup,
		}
	}, func(items interface{}) int {
		return len(items.([]*sql.Color))
	}, nil)

	if err != nil {
		return err
	}

	fmt.Println("Migração de cores concluída")
	return nil
}
