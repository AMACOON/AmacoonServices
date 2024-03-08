package catshowclass

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Code        string `gorm:"type:varchar(10);not null"`
	Name        string `gorm:"type:varchar(10);not null"`
	Description string `gorm:"type:varchar(50);not null"`
	Order       int    `gorm:"not null"`
	NewOrder    int    `gorm:"not null"`
}

func (Class) TableName() string {
	return "cat_show_classes"
}
