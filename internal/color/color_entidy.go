package color

import (
    "gorm.io/gorm"
)

type Color struct {
    gorm.Model
    BreedCode string
    EmsCode   string
    Name      string
    Group     int
    SubGroup  int
}

func (Color) TableName() string {
    return "colors"
}
