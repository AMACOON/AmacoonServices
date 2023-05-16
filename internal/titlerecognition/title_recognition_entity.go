package titlerecognition

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/catservice"
	"gorm.io/gorm"
)

type TitleRecognition struct {
	gorm.Model
	CatData        catservice.CatService   `gorm:"embedded;embeddedPrefix:cat_"`
	OwnerData      catservice.OwnerService `gorm:"embedded;embeddedPrefix:owner_"`
	Titles         []Title            `gorm:"foreignKey:TitleRecognitionID"`
	Status         string
	ProtocolNumber string
	RequesterID    string
}

func (TitleRecognition) TableName() string {
	return "service_title_recognition"
}

type Title struct {
	gorm.Model
	TitleRecognitionID uint // Esta é a chave estrangeira para a tabela title_recognition
	TitleID            uint
	TitleCode          string
	TitleName          string
	Certificate        string
	Date               time.Time
	Judge              string
}

func (Title) TableName() string {
	return "service_title_recognition_titles"
}
