package catshowresult

import (
	"gorm.io/gorm"
)

type CatShowResult struct {
	gorm.Model
	RegistrationID        *uint                `gorm:"not null"`
	CatShowID             *uint                `gorm:"not null"`
	CatShowSubID          *uint                `gorm:"not null"`
	Number                int                  `gorm:"not null"`
	CatShowResultMatrixID *uint                `gorm:"not null"`
	CatShowResultMatrix   *CatShowResultMatrix `gorm:"foreignKey:CatShowResultMatrixID"`
}

func (CatShowResult) TableName() string {
	return "cat_show_results"
}

type CatShowResultMatrix struct {
	gorm.Model
	CatShowID   *uint  `gorm:"not null"`
	Description string `gorm:"not null"`
	Score       int    `gorm:"not null"`
}

func (CatShowResultMatrix) TableName() string {
	return "cat_show_results_ranking_matrix"
}
