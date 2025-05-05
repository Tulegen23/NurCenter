package models

import "time"

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	IsDone      bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
