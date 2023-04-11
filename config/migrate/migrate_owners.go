package migrate

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MigrateOwners(db *gorm.DB, client *mongo.Client) error {
	// Busque todos os registros da tabela "expositores" usando GORM
	var owners []*owner.Owner
	if err := db.Unscoped().Find(&owners).Error; err != nil {
		return err
	}

	// Converta os registros do GORM para o formato do MongoDB
	ownerMongos := make([]interface{}, len(owners))
	for i, o := range owners {
		countryId, err := findCountryIdByCode(client, o.Country)
		if err != nil {
			return err
		}

		ownerMongos[i] = owner.OwnerMongo{
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
	}

	// Insira os registros convertidos na coleção do MongoDB
	_, err := client.Database("amacoon").Collection("owners").InsertMany(context.Background(), ownerMongos)
	if err != nil {
		return err
	}

	return nil
}
