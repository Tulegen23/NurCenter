package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
