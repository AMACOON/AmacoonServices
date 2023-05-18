package owner

import (
	"github.com/scuba13/AmacoonServices/internal/club"
	"github.com/scuba13/AmacoonServices/internal/country"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email        string
	PasswordHash string `gorm:"index" json:"-"`
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
	Clubs        []OwnerClub
}

func (Owner) TableName() string {
	return "owners"
}

type OwnerClub struct {
	gorm.Model
	OwnerID   *uint
	ClubID    *uint
	Club      *club.Club `gorm:"foreignKey:ClubID"`
	Associate bool
	Valid     bool
}

func (OwnerClub) TableName() string {
	return "owners_clubs"
}
