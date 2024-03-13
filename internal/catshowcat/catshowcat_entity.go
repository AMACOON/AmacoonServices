package catshowcat

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"gorm.io/gorm"
)

type CatShowCat struct {
	gorm.Model
	RegistrationID   *uint                   `gorm:"column:registration_id"`
	Name             string                 `gorm:"column:name" validate:"required"`
	NameFull         string                 `gorm:"-"`
	Registration     string                 `gorm:"column:registration;index" validate:"required,notzeroes"`
	RegistrationType string                 `gorm:"column:registrationtype"`
	Microchip        string                 `gorm:"column:microchip"`
	Gender           string                 `gorm:"column:gender;index"`
	Birthdate        time.Time              `gorm:"column:birthdate" validate:"required"`
	Neutered         *bool                  `gorm:"column:neutered" validate:"required"`
	Validated        bool                   `gorm:"column:validated"`
	Observation      string                 `gorm:"column:observation"`
	Fifecat          bool                   `gorm:"column:fifecat"`
	FederationID     *uint                  `gorm:"column:federation_id" validate:"required"`
	Federation       *federation.Federation `gorm:"foreignKey:FederationID"`
	BreedID          *uint                  `gorm:"column:breed_id" validate:"required"`
	Breed            *breed.Breed           `gorm:"foreignKey:BreedID"`
	ColorID          *uint                  `gorm:"column:color_id" validate:"required"`
	Color            *color.Color           `gorm:"foreignKey:ColorID"`
	CatteryID        *uint                  `gorm:"column:cattery_id"`
	Cattery          *cattery.Cattery       `gorm:"foreignKey:CatteryID"`
	OwnerID          *uint                  `gorm:"column:owner_id;index" validate:"required"`
	Owner            *owner.Owner           `gorm:"foreignKey:OwnerID"`
	CountryID        *uint                  `gorm:"column:country_id" validate:"required"`
	Country          *country.Country       `gorm:"foreignKey:CountryID"`
	Titles           *[]TitlesCatShowCat    `gorm:"foreignKey:CatShowCatID"`
	Files            *[]FilesCatShowCat     `gorm:"foreignKey:CatShowCatID"`

	FatherID            *uint        `gorm:"column:father_id"`
	FatherName          string       `gorm:"-"`
	FatherBreedId       *uint        `gorm:"-"`
	FatherBreed         *breed.Breed `gorm:"-"`
	FatherColorID       *uint        `gorm:"-"`
	FatherColor         *color.Color `gorm:"-"`
	FatherNameManual    *string      `json:"-"`
	FatherBreedIDManual *uint        `json:"-"`
	FatherColorIDManual *uint        `json:"-"`

	MotherID            *uint        `gorm:"column:mother_id"`
	MotherName          string       `gorm:"-"`
	MotherBreedID       *uint        `gorm:"-"`
	MotherBreed         *breed.Breed `gorm:"-"`
	MotherColorId       *uint        `gorm:"-"`
	MotherColor         *color.Color `gorm:"-"`
	MotherNameManual    *string      `json:"-"`
	MotherBreedIDManual *uint        `json:"-"`
	MotherColorIDManual *uint        `json:"-"`

	FatherNameTemp string `json:"-"`
	MotherNameTemp string `json:"-"`
	CatIDOld       *uint  `json:"-"`

}

func (CatShowCat) TableName() string {
	return "cat_show_cats"
}

type TitlesCatShowCat struct {
	gorm.Model
	CatShowCatID        uint
	TitleID      uint
	Titles       *title.Title `gorm:"foreignkey:TitleID"`
	Date         time.Time
	FederationID uint `gorm:"foreignkey:FederationID"`
}

func (TitlesCatShowCat) TableName() string {
	return "cat_show_cats_titles"
}

type FilesCatShowCat struct {
	gorm.Model
	CatShowCatID    uint
	FileData utils.Files `gorm:"embedded"`
}

func (FilesCatShowCat) TableName() string {
	return "cat_show_cats_files"
}
