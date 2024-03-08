package catshowregistration

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/scuba13/AmacoonServices/internal/catshowclass"
	"github.com/scuba13/AmacoonServices/internal/judge"
	"github.com/scuba13/AmacoonServices/internal/owner"

	"gorm.io/gorm"
)

// Inscricao representa uma inscrição de gato em exposição.
type Registration struct {
	gorm.Model
	CatShowID        *uint                  `gorm:"not null"`
	CatShow          *catshow.CatShow       `gorm:"foreignKey:CatShowID"`
	CatShowSubID     *uint                  `gorm:"not null"`
	CatShowSub       *catshow.CatShowSub    `gorm:"foreignKey:CatShowSubID"`
	OwnerID          *uint                  `gorm:"not null"`
	Owner            *owner.Owner           `gorm:"foreignKey:OwnerID"`
	CatID            *uint                  `gorm:"not null"`
	CatShowCat       *catshowcat.CatShowCat `gorm:"foreignKey:RegistrationID;references:ID"`
	ClassID          *uint                  `gorm:"not null"`
	Class            *catshowclass.Class    `gorm:"foreignKey:ClassID"`
	JudgeID          *uint
	Judge            *judge.Judge `gorm:"foreignKey:JudgeID"`
	RegistrationDate time.Time    `gorm:"not null"`
	Number           int
	Observations     string `gorm:"type:varchar(255)"`
	Updated          bool
	Active           bool
}

func (Registration) TableName() string {
	return "cat_shows_registration"
}

// Inscricao representa uma inscrição de gato em exposição.
type RegistrationUpdated struct {
	gorm.Model
	RegistrationID *uint     `gorm:"not null"`
	ClassID        *uint     `gorm:"not null"`
	ColorID        *uint     `gorm:"not null"`
	JudgeID        *uint     `gorm:"not null"`
	Birthdate      time.Time `gorm:"not null"`
	Gender         string    `gorm:"not null"`
}

func (RegistrationUpdated) TableName() string {
	return "cat_shows_registration_updated"
}
