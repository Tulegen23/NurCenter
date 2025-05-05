package models

import "time"

type Note struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Title     string `gorm:"not null"`
	Content   string
	Category  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
