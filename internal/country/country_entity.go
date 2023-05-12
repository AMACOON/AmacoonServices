package country

import (
    "gorm.io/gorm"
)

type Country struct {
    gorm.Model
    Code        string
    Name        string
    IsActivated bool
}

func (Country) TableName() string {
    return "countries"
}
