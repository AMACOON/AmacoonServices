package title

import (
    "gorm.io/gorm"
)

type Title struct {
    gorm.Model
    Name        string
	Code        string `gorm:"type:varchar(191);unique"`
    Type        string
    Certificate string
    Amount      int
    Observation string
}

func (Title) TableName() string {
    return "titles"
}