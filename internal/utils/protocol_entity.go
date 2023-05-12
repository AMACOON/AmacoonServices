package utils

import (
	"gorm.io/gorm"
)

type Protocol struct {
	gorm.Model
	Protocol string `gorm:"column:protocol"`
}


func (Protocol) TableName() string {
    return "service_protocols"
}