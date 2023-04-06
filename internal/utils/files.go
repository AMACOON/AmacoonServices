package utils

import (
	"gorm.io/gorm"
	
)

type FilesDB struct {
	gorm.Model
	ID       uint `gorm:"primarykey;autoIncrement"`
	Name     string `gorm:"column:nome_arquivo"`
	Type     string `gorm:"column:tipo_arquivo"`
	Base64   string `gorm:"column:string_arquivo"`
	ProtocolNumber string `gorm:"column:numero_protocolo"`
	ServiceID uint `gorm:"column:id_serviço"`
}
func (F *FilesDB) TableName() string {
	return "pdf_serviços"
}



//////////////////////////////////////////

type Files struct {
	ID       uint
	Name     string 
	Type     string 
	Base64   string 
	ProtocolNumber string 
	ServiceID uint
}

