package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Name      string `json:"name"`
	Streak    int    `json:"streak"`
	Completed bool   `json:"completed"`
}
