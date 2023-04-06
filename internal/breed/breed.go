package breed

import "gorm.io/gorm"

type Breed struct {
	*gorm.Model
	BreedID  string `gorm:"column:id_racas;primaryKey"`
	BreedName     string `gorm:"column:nome"`
	BreedCategory int    `gorm:"column:categoria"`
	BreedByGroup  string `gorm:"column:por_grupo"`
}

func (b *Breed) TableName() string {
	return "racas"
}



type BreedCompatibility struct {
	*gorm.Model
	IDRaca1 string `gorm:"primaryKey;column:id_racas1"`
	IDRaca2 string `gorm:"primaryKey;column:id_racas2"`
}

func (b *BreedCompatibility) TableName() string {
	return "racas_compat"
}
