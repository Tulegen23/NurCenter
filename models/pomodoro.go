package models

import "gorm.io/gorm"

type Pomodoro struct {
	gorm.Model
	SessionType string `json:"session_type"`
	Duration    int    `json:"duration"`
}
