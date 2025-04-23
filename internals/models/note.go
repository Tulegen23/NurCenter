package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"size:200"`
	Content  string
	Category string `gorm:"size:50"`
}
