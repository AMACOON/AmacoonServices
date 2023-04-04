package models

import "gorm.io/gorm"

type Owner struct {
*gorm.Model
OwnerID uint `gorm:"column:id_expositores;primaryKey"`
OwnerName string `gorm:"column:nome"`
Address string `gorm:"column:endereco"`
Complement string `gorm:"column:complemento"`
Neighborhood string `gorm:"column:bairro"`
ZipCode string `gorm:"column:cep"`
City string `gorm:"column:cidade"`
State string `gorm:"column:estado"`
Phone string `gorm:"column:telefone"`
}

func (o *Owner) TableName() string {
return "expositores"
}