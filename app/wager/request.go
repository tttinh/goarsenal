package wager

type CreateWagerRequest struct {
	TotalWagerValue   uint `json:"total_wager_value"`
	Odds              uint `json:"odds"`
	SellingPercentage uint `json:"selling_percentage"`
	SellingPrice      uint `json:"selling_price"`
}
