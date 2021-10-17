package entity

import (
	"time"
)

type Wager struct {
	ID                int
	TotalWagerValue   int
	Odds              int
	SellingPercentage int
	SellingPrice      int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// TableName sets the table name for the entity.
func (Wager) TableName() string {
	return "wager"
}
