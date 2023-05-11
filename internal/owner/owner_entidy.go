package owner

import (
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/internal/country"
)

type Owner struct {
	gorm.Model
	Email        string
	PasswordHash string
	Name         string
	CPF          string
	Address      string
	City         string
	State        string
	ZipCode      string
	CountryID    *uint 
	Country      *country.Country `gorm:"foreignKey:CountryID"`
	Phone        string
	Valid        bool
	ValidId      string
	Observation  string
}

func (Owner) TableName() string {
	return "owners"
}
