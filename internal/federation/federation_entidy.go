package federation

import (
	
	"gorm.io/gorm"
)

type Federation struct {
	gorm.Model
	Name           string          
	FederationCode string          
	CountryID        *uint                  `gorm:"foreignKey:CountryID"`    
	
}

func (Federation) TableName() string {
	return "federations"
}
