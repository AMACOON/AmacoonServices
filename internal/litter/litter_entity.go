package litter

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/catservice"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"gorm.io/gorm"
)

type Litter struct {
	gorm.Model
	MotherData     catservice.CatService   `gorm:"embedded;embeddedPrefix:mother_"`
	FatherData     catservice.CatService   `gorm:"embedded;embeddedPrefix:father_"`
	MotherOwner    catservice.OwnerService `gorm:"embedded;embeddedPrefix:motherOwner_"`
	FatherOwner    catservice.OwnerService `gorm:"embedded;embeddedPrefix:fatherOwner_"`
	CatteryName    string
	NumKittens     int
	BirthDate      time.Time
	CountryCode    string
	Status         string
	ProtocolNumber string
	RequesterID    uint
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
	BreedName  string
	ColorName  string
	EmsCode    string
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
	LitterID  uint 
	FileData utils.Files `gorm:"embedded"`
}

func (FilesLitter) TableName() string {
	return "service_litters_files"
}
