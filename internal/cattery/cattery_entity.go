package cattery

import (
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"gorm.io/gorm"
)

type Cattery struct {
	gorm.Model
	Name        string
	BreederName string
	OwnerID     *uint
	Owner       *owner.Owner `gorm:"foreignKey:OwnerID"`
	CountryID   *uint
	Country     *country.Country `gorm:"foreignKey:CountryID"`
}

func (Cattery) TableName() string {
	return "catteries"
}
