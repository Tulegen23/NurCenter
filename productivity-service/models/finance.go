package models

import "time"

type Finance struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Description string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
