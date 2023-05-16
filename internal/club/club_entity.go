package club

import (
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name string
	Nickname string
	Email string
	Login string
	Password string
	Permission int
	}

func (Club) TableName() string {
	return "clubs"
}
