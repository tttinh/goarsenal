package entity

import (
	"time"
)

type Purchase struct {
	ID          uint32 `gorm:"primaryKey"`
	BuyingPrice float32
	BoughtAt    time.Time `gorm:"autoCreateTime"`

	WagerID uint32
	Wager   Wager
}

// TableName sets the table name for the entity.
func (Purchase) TableName() string {
	return "purchase"
}
