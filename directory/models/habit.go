package models

import (
	"gorm.io/gorm"
	"time"
)

type Habit struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Name   string `gorm:"size:100"`
	Logs   []HabitLog
}

type HabitLog struct {
	gorm.Model
	HabitID uint      `gorm:"not null"`
	LogDate time.Time `gorm:"type:date"`
}
