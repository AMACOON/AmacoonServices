package catshow

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/club"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/judge"
	"gorm.io/gorm"
)

type CatShow struct {
	gorm.Model
	FederationID        *uint                  `gorm:"column:federation_id;not null" validate:"required"`
	Federation          *federation.Federation `gorm:"foreignKey:FederationID;not null"`
	ClubID              *uint                  `gorm:"column:club_id;not null" validate:"required"`
	Club                *club.Club             `gorm:"foreignKey:ClubID;not null"`
	Description         string                 `gorm:"type:varchar(255);not null" validate:"required"`
	Location            string                 `gorm:"type:varchar(255);not null" validate:"required"`
	City                string                 `gorm:"type:varchar(255);not null" validate:"required"`
	State               string                 `gorm:"type:varchar(255);not null" validate:"required"`
	CountryID           *uint                  `gorm:"not null" validate:"required"`
	Country             *country.Country       `gorm:"foreignKey:CountryID"`
	StartDate           time.Time              `gorm:"not null" validate:"required"`
	EndDate             time.Time              `gorm:"not null" validate:"required"`
	Finished            bool                   `gorm:"not null" validate:"required"` // false por padrão
	RegistrationStart   time.Time              `gorm:"not null" validate:"required"`
	RegistrationEnd     time.Time              `gorm:"not null" validate:"required"`
	Separated           bool                   `gorm:"not null" validate:"required"` // false por padrão
	MaxCats             int                    `gorm:"not null" validate:"required"`
	MaxCatsPerExhibitor int                    `gorm:"not null" validate:"required"`
	Validated           bool                   `gorm:"not null" validate:"required"` // false por padrão
	CatsDesignated      bool                   `gorm:"not null" validate:"required"` // false por padrão
	Certificate         string                 `gorm:"type:varchar(255);not null" validate:"required"`
	DatesDescription    string                 `gorm:"type:varchar(255);not null" validate:"required"`
	CatShowSubs         []CatShowSub           `gorm:"foreignKey:CatShowID"`
	CatShowJudges       []CatShowJudge         `gorm:"foreignKey:CatShowID"`
}

func (CatShow) TableName() string {
	return "cat_shows"
}

type CatShowSub struct {
	gorm.Model
	CatShowID     uint      `gorm:"not null;foreignKey:CatShowID" validate:"required"` // foreign key
	CatShowNumber int       `gorm:"not null" validate:"required"`
	Description   string    `gorm:"type:varchar(120);not null" validate:"required"`
	CatShowDate   time.Time `gorm:"not null" validate:"required"`
	CatShowType   string    `gorm:"type:varchar(1);not null" validate:"required"`
}

func (CatShowSub) TableName() string {
	return "cat_show_subs"
}

type CatShowJudge struct {
	gorm.Model
	CatShowID uint        `gorm:"not null;foreignKey:CatShowID" validate:"required"` // foreign key
	JudgeID   uint        `gorm:"not null" validate:"required"`
	Judge     judge.Judge `gorm:"foreignKey:JudgeID"`
}

func (CatShowJudge) TableName() string {
	return "cat_show_judges"
}
