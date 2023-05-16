package federation

import (
	
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/internal/country"
)

type Federation struct {
	gorm.Model
	Name           string          
	FederationCode string          
	CountryID        *uint                  
	Country 	   *country.Country `gorm:"foreignKey:CountryID"`  
	
}

func (Federation) TableName() string {
	return "federations"
}
