package models

import "gorm.io/gorm"

type Country struct {
	*gorm.Model
	CountryCode      string `gorm:"primaryKey;column:code"`
	CountryName string `gorm:"column:descricao"`
	Activate   string `gorm:"column:visivel"`
}

func (c *Country) TableName() string {
	return "country_codes"
}
