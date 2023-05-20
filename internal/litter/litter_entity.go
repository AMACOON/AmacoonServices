package litter

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/catservice"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"gorm.io/gorm"
)

type Litter struct {
	gorm.Model
	MotherData     catservice.CatServiceEntity `gorm:"embedded;embeddedPrefix:mother_"`
	FatherData     catservice.CatServiceEntity `gorm:"embedded;embeddedPrefix:father_"`
	NumKittens     int
	BirthDate      time.Time
	Status         string
	ProtocolNumber string
	RequesterID    uint
	CatteryID      uint
	CountryID      uint
	KittenData     *[]KittenLitter `gorm:"foreignKey:LitterID"`
	Files          *[]FilesLitter  `gorm:"foreignKey:LitterID"`
}

func (Litter) TableName() string {
	return "service_litters"
}

type KittenLitter struct {
	gorm.Model
	Name       string
	Gender     string
	BreedID    uint
	ColorID    uint
	ColorNameX string
	Microchip  string
	Breeding   bool
	LitterID   uint
}

func (KittenLitter) TableName() string {
	return "service_litters_kittens"
}

type FilesLitter struct {
	gorm.Model
	LitterID uint
	FileData utils.Files `gorm:"embedded"`
}

func (FilesLitter) TableName() string {
	return "service_litters_files"
}
