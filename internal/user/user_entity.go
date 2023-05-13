package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OwnerID      *uint `gorm:"foreignKey:OwnerID"`
	Email        string `gorm:"index"`
	PasswordHash string `gorm:"index" json:"-"`
	Name         string
	CPF          string
	IsAdmin      bool
}

func (User) TableName() string {
	return "users"
}

