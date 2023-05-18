package judge

import (
	"github.com/scuba13/AmacoonServices/internal/country"
	"gorm.io/gorm"
)

type Judge struct {
	gorm.Model
	Name      string
	Email     string
	CountryID *uint
	Country   *country.Country `gorm:"foreignKey:CountryID"`
	Category1A bool
	Category1B bool
	Category2 bool
	Category3 bool
	Category4C bool
	Category4D bool
}

func (Judge) TableName() string {
	return "judges"
}
