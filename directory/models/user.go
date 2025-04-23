package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null;size:50"`
	Email     string `gorm:"unique;not null;size:100"`
	Password  string `gorm:"not null"`
	Todos     []Todo
	Notes     []Note
	Habits    []Habit
	Goals     []Goal
	Finances  []Finance
	Reminders []Reminder
}
