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
	Category1 bool
	Category2 bool
	Category3 bool
	Category4 bool
}

func (Judge) TableName() string {
	return "judges"
}
