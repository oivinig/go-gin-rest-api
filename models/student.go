package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentName string `json:"student_name"`
	Document    string `json:"document"`
	Registry    string `json:"registry"`
}
