package wager

type CreateWagerRequest struct {
	TotalWagerValue   int `json:"total_wager_value"`
	Odds              int `json:"odds"`
	SellingPercentage int `json:"selling_percentage"`
	SellingPrice      int `json:"selling_price"`
}
