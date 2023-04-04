package models

import (
	"gorm.io/gorm"
	"time"
)

type Protocol struct {
	ID            uint           `gorm:"primary_key"`
	ProtocolNumber string         `gorm:"unique;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (P *Protocol) TableName() string {
	return "protocolos"
}
