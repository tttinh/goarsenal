package entity

import (
	"time"
)

type Wager struct {
	ID                  uint `gorm:"primaryKey"`
	TotalWagerValue     uint
	Odds                uint
	SellingPercentage   uint
	SellingPrice        uint
	CurrentSellingPrice uint
	PercentageSold      uint
	AmountSold          uint
	PlacedAt            time.Time `gorm:"autoCreateTime"`
}

// TableName sets the table name for the entity.
func (Wager) TableName() string {
	return "wager"
}
