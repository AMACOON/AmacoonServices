package models

import (
	"gorm.io/gorm"
	"time"
)
type Owner struct {
    gorm.Model
    OwnerID      uint   `gorm:"column:id_expositores;primaryKey"`
    Email        string `gorm:"column:email;unique"`
    PasswordHash string `gorm:"column:senha"`
    OwnerName    string `gorm:"column:nome"`
    Address      string `gorm:"column:endereco"`
    City         string `gorm:"column:cidade"`
    State        string `gorm:"column:estado"`
    ZipCode      string `gorm:"column:cep"`
    Country      string `gorm:"column:pais"`
    Phone        string `gorm:"column:telefone"`
    Valid        string `gorm:"column:valido"`
    ValidationID string `gorm:"column:id_validacao"`
    Observation  []byte `gorm:"column:observacao"`
    CreatedAt    time.Time `gorm:"column:datacadastro"`
    CPF          string `gorm:"column:cpf;default:0"`
}

func (o *Owner) TableName() string {
    return "expositores"
}
