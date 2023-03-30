package models

import "gorm.io/gorm"

type Litter struct {
	gorm.Model
	LitterID      uint   `gorm:"column:id_ninhadas;primaryKey;autoIncrement"`
	FatherName    string `gorm:"column:nome_pai"`
	FatherReg     string `gorm:"column:registro_pai"`
	FatherMicro   string `gorm:"column:microchip_pai"`
	FatherBreed   string `gorm:"column:raça_pai"`
	FatherEMSCode string `gorm:"column:id_cor_pai"`
	FatherColor   string `gorm:"column:cor_pai"`
	FatherOwnerID uint   `gorm:"column:id_expositores_pai"`
	MotherName    string `gorm:"column:nome_mae"`
	MotherReg     string `gorm:"column:registro_mae"`
	MotherMicro   string `gorm:"column:microchip_mae"`
	MotherBreed   string `gorm:"column:raça_mae"`
	MotherEMSCode string `gorm:"column:id_cor_mae"`
	MotherColor   string `gorm:"column:cor_mae"`
	MotherOwnerID uint   `gorm:"column:id_expositores_mae"`
	CatteryID     uint `gorm:"column:id_gatis"`
	CatteryName   string `gorm:"column:nome_gatil"`
	NumKittens    int   `gorm:"column:num_filhotes"`
	BirthDate     string `gorm:"column:data_nascimento"`
	Status        string `gorm:"column:status"`
}

func (l *Litter) TableName() string {
	return "ninhadas"
}

type Kitten struct {
	gorm.Model
	KittenID    uint   `gorm:"column:id_filhotes;primaryKey;autoIncrement"`
	LitterID    uint   `gorm:"column:id_ninhadas"`
	BreedName   string `gorm:"column:nome_raça"`
	ColorName   string `gorm:"column:nome_cor"`
	EmsCodeID   string `gorm:"column:emscode"`
	CountryCode string `gorm:"column:codigo_pais"`
	Microchip   string `gorm:"column:microchip"`
	ColorNameX  string `gorm:"column:nome_cor_x"`
	Breeding    bool   `gorm:"column:reprodução"`
	Name        string `gorm:"column:nome__gato"`
	Sex         string `gorm:"column:sexo"`
	Status      string `gorm:"column:status"`
}

func (k *Kitten) TableName() string {
	return "filhotes"
}

//////////////////////////////
// Models para Serviço

type LitterData struct {
	MotherData  CatData
	FatherData  CatData
	BirthData   BirthData
	KittenData  []Kitten
	LitterID    uint
	Status      string
}

type CatData struct {
	Name    string
	Registration     string 
	Microchip   string 
	BreedName   string 
	EmsCode string 
	ColorName   string 
	OwnerID uint
}

type BirthData struct {
	CatteryID     uint 
	CatteryName   string 
	NumKittens    int   
	BirthDate     string
}
