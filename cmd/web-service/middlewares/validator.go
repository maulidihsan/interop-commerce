package middlewares

import (
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func ValidatorInit() {
	validate = validator.New()
}

func GetValidator() *validator.Validate {
	return validate
}