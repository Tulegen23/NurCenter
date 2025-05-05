package models

import "time"

type Habit struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Frequency string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
