package membership

import (
	"time"

	"gorm.io/gorm"
)

type MembershipRequest struct {
	gorm.Model

	ProtocolCode    string `gorm:"uniqueIndex"`
	AssociationType string // individual | family

	// Dados do titular
	OwnerName     string
	OwnerLastName string
	CPF           string
	Email         string
	Phone         string
	Address       string
	City          string
	State         string
	ZipCode       string

	// Observações
	HasFifeCattery bool
	Observation    string

	Status string // PENDING | APPROVED | REJECTED

	Cats  []MembershipCat  `gorm:"foreignKey:MembershipID"`
	Files []MembershipFile `gorm:"foreignKey:MembershipID"`

	SubmittedAt time.Time
}

func (MembershipRequest) TableName() string {
	return "membership_requests"
}

type MembershipCat struct {
	gorm.Model
	MembershipID uint

	Name      string
	BirthDate string
	Sex       string
	BreedID   uint
}

func (MembershipCat) TableName() string {
	return "membership_cats"
}

type MembershipFile struct {
	gorm.Model
	MembershipID uint

	FileType string // owner_document | cat_document | other
	S3Key    string
	FileName string
}

func (MembershipFile) TableName() string {
	return "membership_files"
}
