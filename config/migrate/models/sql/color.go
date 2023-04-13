package sql

import "gorm.io/gorm"

type Color struct {
	*gorm.Model
	ColorID   int    `gorm:"column:id_cores;primaryKey"`
	BreedID   string `gorm:"column:id_raca"`
	EmsCode string `gorm:"column:id_emscode"`
	ColorName      string `gorm:"column:descricao"`
	Group     int    `gorm:"column:grupo"`
	SubGroup  int    `gorm:"column:sub_grupo"`
}

func (c *Color) TableName() string {
	return "cores"
}
