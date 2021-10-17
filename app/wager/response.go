package wager

type WagerResponse struct {
	ID                  uint   `json:"id"`
	TotalWagerValue     uint   `json:"total_wager_value"`
	Odds                uint   `json:"odds"`
	SellingPercentage   uint   `json:"selling_percentage"`
	SellingPrice        uint   `json:"selling_price"`
	CurrentSellingPrice uint   `json:"current_selling_price"`
	PercentageSold      uint   `json:"percentage_sold"`
	AmountSold          uint   `json:"amount_sold"`
	PlacedAt            uint64 `json:"placed_at"`
}
