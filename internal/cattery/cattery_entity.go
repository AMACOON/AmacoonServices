package cattery

import (
	
	"gorm.io/gorm"
)

type Cattery struct {
	gorm.Model
	Name        string
	BreederName string
	OwnerID          *uint                  `gorm:"foreignKey:OwnerID"`
	CountryID        *uint                  `gorm:"foreignKey:CountryID"`
}

func (Cattery) TableName() string {
    return "catteries"
}

