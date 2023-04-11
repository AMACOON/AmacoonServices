package cattery

import "gorm.io/gorm"

type Cattery struct {
	*gorm.Model
	BreederID  string `gorm:"column:id_gatis;primaryKey"`
	BreederName     string `gorm:"column:nome_gatil"`
	BreederOwner string    `gorm:"column:criador_gatil"`
	BreederCountry  string `gorm:"column:pais_gatil"`
}

func (b *Cattery) TableName() string {
	return "gatis"
}