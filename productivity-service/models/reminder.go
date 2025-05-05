package models

import "time"

type Reminder struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"not null"`
	DueDate   time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
