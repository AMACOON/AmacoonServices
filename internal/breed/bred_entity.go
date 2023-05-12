package breed

import (
	"gorm.io/gorm"
)

type Breed struct {
	gorm.Model
	BreedCode     string `gorm:"type:varchar(191);unique"`
	BreedName     string
	BreedCategory int
	BreedByGroup  string
}

func (Breed) TableName() string {
	return "breeds"
}

type BreedCompatibility struct {
	gorm.Model
	BreedCode1 string
	BreedCode2 string
}

func (BreedCompatibility) TableName() string {
	return "breed_compatibilities"
}
