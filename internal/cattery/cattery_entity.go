package cattery

import (
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/internal/utils"
)

type Cattery struct {
	gorm.Model
	Name        string
	BreederName string
	OwnerID     *uint
	Owner       *owner.Owner `gorm:"foreignKey:OwnerID"`
	CountryID   *uint
	Country     *country.Country `gorm:"foreignKey:CountryID"`
	Files       []FilesCattery `gorm:"foreignKey:CatteryID"`
}

func (Cattery) TableName() string {
	return "catteries"
}

type FilesCattery struct {
	gorm.Model
	CatteryID    uint
	FileData utils.Files `gorm:"embedded"`
}

func (FilesCattery) TableName() string {
	return "catteries_files"
}

