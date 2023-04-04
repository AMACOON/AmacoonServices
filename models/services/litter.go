package models

import "gorm.io/gorm"

type Litter struct {
	gorm.Model
	FatherName      string `gorm:"column:nome_pai"`
	FatherReg       string `gorm:"column:registro_pai"`
	FatherMicro     string `gorm:"column:microchip_pai"`
	FatherBreed     string `gorm:"column:raça_pai"`
	FatherEMSCode   string `gorm:"column:id_cor_pai"`
	FatherColor     string `gorm:"column:cor_pai"`
	FatherOwnerID   uint   `gorm:"column:id_expositores_pai"`
	FatherOwnerName string `gorm:"column:nome_expositor_pai"`
	FatherAddress   string `gorm:"column:endereço_expositor_pai"`
	FatherZipCode   string `gorm:"column:cep_expositor_pai"`
	FatherCity      string `gorm:"column:cidade_expositor_pai"`
	FatherState     string `gorm:"column:estado_expositor_pai"`
	FatherCountry   string `gorm:"column:pais_expositor_pai"`
	FatherPhone     string `gorm:"column:telefone_expositor_pai"`

	MotherName      string `gorm:"column:nome_mae"`
	MotherReg       string `gorm:"column:registro_mae"`
	MotherMicro     string `gorm:"column:microchip_mae"`
	MotherBreed     string `gorm:"column:raça_mae"`
	MotherEMSCode   string `gorm:"column:id_cor_mae"`
	MotherColor     string `gorm:"column:cor_mae"`
	MotherOwnerID   uint   `gorm:"column:id_expositores_mae"`
	MotherOwnerName string `gorm:"column:nome_expositor_mae"`
	MotherAddress   string `gorm:"column:endereço_expositor_mae"`
	MotherZipCode   string `gorm:"column:cep_expositor_mae"`
	MotherCity      string `gorm:"column:cidade_expositor_mae"`
	MotherState     string `gorm:"column:estado_expositor_mae"`
	MotherCountry   string `gorm:"column:pais_expositor_mae"`
	MotherPhone     string `gorm:"column:telefone_expositor_mae"`

	CatteryID   uint   `gorm:"column:id_gatis"`
	CatteryName string `gorm:"column:nome_gatil"`
	NumKittens  int    `gorm:"column:num_filhotes"`
	BirthDate   string `gorm:"column:data_nascimento"`
	Country     string `gorm:"column:pais_ninhada"`
	ProtocolNumber string `gorm:"not null"`
	Status string `gorm:"column:status"`
}

func (l *Litter) TableName() string {
	return "ninhadas"
}

type Kitten struct {
	gorm.Model
	LitterID   uint   `gorm:"column:id_ninhadas"`
	BreedName  string `gorm:"column:nome_raça"`
	ColorName  string `gorm:"column:nome_cor"`
	EmsCodeID  string `gorm:"column:emscode"`
	Microchip  string `gorm:"column:microchip"`
	ColorNameX string `gorm:"column:nome_cor_x"`
	Breeding   bool   `gorm:"column:reprodução"`
	Name       string `gorm:"column:nome__gato"`
	Sex        string `gorm:"column:sexo"`
	Status     string `gorm:"column:status"`
}

func (k *Kitten) TableName() string {
	return "filhotes"
}

//////////////////////////////
// Models para Serviço

type LitterData struct {
	MotherData CatData
	FatherData CatData
	BirthData  BirthData
	KittenData []KittenService
	LitterID   uint
	Status     string
}

type CatData struct {
	Name         string
	Registration string
	Microchip    string
	BreedName    string
	EmsCode      string
	ColorName    string
	OwnerID      uint
	OwnerName    string
	Address      string
	ZipCode      string
	City         string
	State        string
	Country      string
	Phone        string
}

type BirthData struct {
	CatteryID   uint
	CatteryName string
	NumKittens  int
	BirthDate   string
	Country     string
}

type KittenService struct {
	KittenID   *uint
    LitterID   *uint
	BreedName  string
	ColorName  string
	EmsCodeID  string
	Microchip  string
	ColorNameX string
	Breeding   bool
	Name       string
	Sex        string
	Status     string
}
