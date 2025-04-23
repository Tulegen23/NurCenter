package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Finance struct {
	gorm.Model
	UserID      uint            `gorm:"not null"`
	Amount      decimal.Decimal `gorm:"type:numeric(10,2);not null"`
	Type        string          `gorm:"size:10;check:type IN ('income', 'expense')"`
	Category    string          `gorm:"size:100"`
	Description string
	Date        time.Time `gorm:"type:date;default:CURRENT_DATE"`
}
