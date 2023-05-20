package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func NotZeroes(fl validator.FieldLevel) bool {
	if fl.Field().String() == "000000" {
		return false
	}
	return true
}

func init() {
	Validate = validator.New()
	_ = Validate.RegisterValidation("notzeroes", NotZeroes)
}

func ValidateStruct(s interface{}) error {
	return Validate.Struct(s)
}
