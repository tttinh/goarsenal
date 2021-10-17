package wager

type WagerResponse struct {
	ID                  uint32  `json:"id"`
	TotalWagerValue     uint32  `json:"total_wager_value"`
	Odds                uint32  `json:"odds"`
	SellingPercentage   uint32  `json:"selling_percentage"`
	SellingPrice        float32 `json:"selling_price"`
	CurrentSellingPrice float32 `json:"current_selling_price"`
	PercentageSold      *uint32 `json:"percentage_sold"`
	AmountSold          *uint32 `json:"amount_sold"`
	PlacedAt            uint64  `json:"placed_at"`
}
