package models

import (
	"gorm.io/gorm"
	"time"
)

type Reminder struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"size:255"`
	Message  string
	RemindAt time.Time
	IsSent   bool `gorm:"default:false"`
}
