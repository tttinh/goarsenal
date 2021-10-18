package wager

type CreateWagerRequest struct {
	TotalWagerValue   uint32  `json:"total_wager_value" binding:"gt=0"`
	Odds              uint32  `json:"odds" binding:"gt=0"`
	SellingPercentage uint32  `json:"selling_percentage" binding:"gte=1,lte=100"`
	SellingPrice      float32 `json:"selling_price" binding:"gt=0"`
}

type BuyWagerRequest struct {
	BuyingPrice float32 `json:"buying_price" binding:"gt=0"`
}
