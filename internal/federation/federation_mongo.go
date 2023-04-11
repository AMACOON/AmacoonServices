package federation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"

)

type FederationMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	FederationCode string             `bson:"federationCode"`
	CountryId      primitive.ObjectID `bson:"countryId"`
}

type Federation struct {
	*gorm.Model
	ID      string `gorm:"primaryKey;column:id_federacoes"`
	FederationCode string `gorm:"column:sigla_federacoes"`
	FederationName   string `gorm:"column:descricao"`
	CountryCode   string `gorm:"column:country_code"`
}

func (c *Federation) TableName() string {
	return "federacoes"
}