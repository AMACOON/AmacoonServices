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
	PasswordHash string
	Permission int
	}

func (Club) TableName() string {
	return "clubs"
}
