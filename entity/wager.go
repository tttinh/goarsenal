package entity

import (
	"time"
)

type Wager struct {
	ID                  uint32 `gorm:"primaryKey"`
	TotalWagerValue     uint32
	Odds                uint32
	SellingPercentage   uint32
	SellingPrice        float32
	CurrentSellingPrice float32
	PercentageSold      *uint32
	AmountSold          *uint32
	PlacedAt            time.Time `gorm:"autoCreateTime"`
}

// TableName sets the table name for the entity.
func (Wager) TableName() string {
	return "wager"
}
