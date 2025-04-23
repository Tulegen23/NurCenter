package models

import (
	"gorm.io/gorm"
	"time"
)

type Goal struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"size:200"`
	Description string
	Deadline    time.Time `gorm:"type:date"`
	IsCompleted bool      `gorm:"default:false"`
	Subgoals    []Subgoal
}

type Subgoal struct {
	gorm.Model
	GoalID uint   `gorm:"not null"`
	Title  string `gorm:"size:200"`
	IsDone bool   `gorm:"default:false"`
}
