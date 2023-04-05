package models

import "gorm.io/gorm"

type TransferDB struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	CatID             int    `gorm:"column:id_gato"`
	CatName           string `gorm:"column:nome_gato"`
	CatBreeding       bool   `gorm:"column:criacao_gato"`
	CatRegistration   string `gorm:"column:registro_gato"`
	CatPedigree       string `gorm:"column:pedigree_gato"`
	CatMicrochip      string `gorm:"column:microchip_gato"`
	CatColorID        int    `gorm:"column:id_cor_gato"`
	CatColorName      string `gorm:"column:nome_cor_gato"`
	CatEmsCode        string `gorm:"column:id_emscode_gato"`
	CatSex            string `gorm:"column:sexo_gato"`
	CatBirthdate      string `gorm:"column:nascimento_gato"`
	CatCountryCode    string `gorm:"column:pais_gato_gato"`
	CatFatherName     string `gorm:"column:nome_pai_gato"`
	CatMotherName     string `gorm:"column:nome_mae_gato"`
	SellerName        string `gorm:"column:nome_vendedor"`
	SellerEmail       string `gorm:"column:email_vendedor"`
	SellerPhone       string `gorm:"column:telefone_vendedor"`
	SellerMobilePhone string `gorm:"column:celular_vendedor"`
	SellerZipCode     string `gorm:"column:cep_vendedor"`
	SellerAddress     string `gorm:"column:endereco_vendedor"`
	SellerDistrict    string `gorm:"column:bairro_vendedor"`
	SellerCity        string `gorm:"column:cidade_vendedor"`
	SellerNumber      string `gorm:"column:numero_vendedor"`
	SellerComplement  string `gorm:"column:complemento_vendedor"`
	SellerCountry     string `gorm:"column:pais_vendedor"`
	BuyerName         string `gorm:"column:nome_comprador"`
	BuyerDocument     string `gorm:"column:documento_comprador"`
	BuyerEmail        string `gorm:"column:email_comprador"`
	BuyerPhone        string `gorm:"column:telefone_comprador"`
	BuyerMobilePhone  string `gorm:"column:celular_comprador"`
	BuyerZipCode      string `gorm:"column:cep_comprador"`
	BuyerAddress      string `gorm:"column:endereco_comprador"`
	BuyerDistrict     string `gorm:"column:bairro_comprador"`
	BuyerCity         string `gorm:"column:cidade_comprador"`
	BuyerNumber       string `gorm:"column:numero_comprador"`
	BuyerComplement   string `gorm:"column:complemento_comprador"`
	BuyerCountry      string `gorm:"column:pais_comprador"`
	Status            string `gorm:"column:status"`
	ProtocolNumber    string `gorm:"not null"`
}

func (c *TransferDB) TableName() string {
	return "transferencia"
}

//////////////////////////////////////////////////////////////////////

//Modelagem Serviço

type CatTransfer struct {
	ID           uint
	CatID        int
	Name         string
	Breeding     bool
	Registration string
	Pedigree     string
	Microchip    string
	ColorID      int
	ColorName    string
	EmsCode      string
	Sex          string
	Birthdate    string
	CountryCode  string
	FatherName   string
	MotherName   string
}

type SellerTransfer struct {
	Name         string
	Email        string
	Phone        string
	MobilePhone  string
	Cep          string
	Address      string
	Neighborhood string
	City         string
	Number       string
	Complement   string
	Country      string
}

type BuyerTransfer struct {
	Name         string
	CpfOrDoc     string
	Email        string
	Phone        string
	MobilePhone  string
	Cep          string
	Address      string
	Neighborhood string
	City         string
	Number       string
	Complement   string
	Country      string
}

type Transfer struct {
	ID             uint
	Cat            CatTransfer
	Seller         SellerTransfer
	Buyer          BuyerTransfer
	Status         string
	ProtocolNumber string
}
