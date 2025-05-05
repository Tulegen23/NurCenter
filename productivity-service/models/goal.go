package models

import "time"

type Goal struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Title     string `gorm:"not null"`
	Deadline  time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
