package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Student struct {
	Name   string `json:"name" validates:"required"`
	RollNo int    `json:"rollNo" validates:"required,RollNo"`
}

func (s *Student) Validates() error {
	validate := validator.New()
	validate.RegisterValidation("RollNo", ValidateRollno)
	return validate.Struct(s)
}

func ValidateRollno(level validator.FieldLevel) bool {
	if level.Field().Int() <= 0 {
		return false
	}
	return true

}

func main() {
	fmt.Println()
}
