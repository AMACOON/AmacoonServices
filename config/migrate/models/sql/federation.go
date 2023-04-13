package sql

import (
	"gorm.io/gorm"
)

type Federation struct {
	*gorm.Model
	ID             string `gorm:"primaryKey;column:id_federacoes"`
	FederationCode string `gorm:"column:sigla_federacoes"`
	FederationName string `gorm:"column:descricao"`
	CountryCode    string `gorm:"column:country_code"`
}

func (c *Federation) TableName() string {
	return "federacoes"
}
