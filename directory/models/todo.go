package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"size:255"`
	Description string
	IsDone      bool `gorm:"default:false"`
}
