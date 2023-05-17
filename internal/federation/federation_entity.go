package federation

import (
	"github.com/scuba13/AmacoonServices/internal/country"
	"gorm.io/gorm"
)

type Federation struct {
	gorm.Model
	Name           string
	FederationCode string
	CountryID      *uint
	Country        *country.Country `gorm:"foreignKey:CountryID"`
}

func (Federation) TableName() string {
	return "federations"
}
