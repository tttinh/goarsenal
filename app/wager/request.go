package wager

type CreateWagerRequest struct {
	TotalWagerValue   uint32  `json:"total_wager_value"`
	Odds              uint32  `json:"odds"`
	SellingPercentage uint32  `json:"selling_percentage"`
	SellingPrice      float32 `json:"selling_price"`
}

type BuyWagerRequest struct {
	BuyingPrice float32 `json:"buying_price"`
}
