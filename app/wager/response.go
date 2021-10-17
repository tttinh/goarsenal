package wager

type WagerResponse struct {
	ID                string `json:"id"`
	TotalWagerValue   int32  `json:"total_wager_value"`
	Odds              int32  `json:"odds"`
	SellingPercentage int32  `json:"selling_percentage"`
	SellingPrice      int32  `json:"selling_price"`
}
