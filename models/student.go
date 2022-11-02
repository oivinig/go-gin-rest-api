package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentName string `json:"student_name" validate:"nonzero"`
	Document    string `json:"document" validate:"len=11, regexp=^[0-9]*$"`
	Registry    string `json:"registry" validate:"len=9, regexp=^[0-11]*$"`
}

func ValidadeStudent(s *Student) error {
	if err := validator.Validate(s); err != nil {
		return err
	}
	return nil
}
