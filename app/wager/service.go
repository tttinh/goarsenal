package wager

import (
	"errors"

	"github.com/tttinh/goarsenal/entity"
	"github.com/tttinh/goarsenal/repository"
)

type serviceImpl struct {
	wagerRepository    repository.WagerRepository
	purchaseRepository repository.PurchaseRepository
}

func (s *serviceImpl) CreateWager(req CreateWagerRequest) (*WagerResponse, error) {
	threshold := float32(req.TotalWagerValue) * (float32(req.SellingPercentage) / 100)
	if req.SellingPrice <= threshold {
		return nil, errors.New("the selling_price is too low")
	}

	wager := &entity.Wager{
		TotalWagerValue:     req.TotalWagerValue,
		Odds:                req.Odds,
		SellingPercentage:   req.SellingPercentage,
		SellingPrice:        req.SellingPrice,
		CurrentSellingPrice: req.SellingPrice,
	}

	if err := s.wagerRepository.Create(wager); err != nil {
		return nil, err
	}

	return &WagerResponse{
		ID:                  wager.ID,
		TotalWagerValue:     wager.TotalWagerValue,
		Odds:                wager.Odds,
		SellingPercentage:   wager.SellingPercentage,
		SellingPrice:        wager.SellingPrice,
		CurrentSellingPrice: wager.CurrentSellingPrice,
		PercentageSold:      wager.PercentageSold,
		AmountSold:          wager.AmountSold,
		PlacedAt:            uint64(wager.PlacedAt.Unix()),
	}, nil
}

func (s *serviceImpl) BuyWager(wagerID uint32, req BuyWagerRequest) (*BuyWagerResponse, error) {
	wager, err := s.wagerRepository.FindWagerByID(wagerID)
	if err != nil {
		return nil, err
	}

	if req.BuyingPrice > wager.CurrentSellingPrice {
		return nil, errors.New("the buying_price is too high")
	}

	if wager.AmountSold == nil {
		wager.AmountSold = &req.BuyingPrice
	} else {
		*wager.AmountSold += req.BuyingPrice
	}

	percentageSold := (*wager.AmountSold) / float32(wager.TotalWagerValue) * 100
	wager.PercentageSold = &percentageSold
	wager.CurrentSellingPrice = float32(wager.TotalWagerValue) - (*wager.AmountSold)

	purchase := &entity.Purchase{
		BuyingPrice: req.BuyingPrice,
		WagerID:     wagerID,
		Wager:       wager,
	}

	if err := s.purchaseRepository.Create(purchase); err != nil {
		return nil, err
	}

	if err := s.wagerRepository.Update(&wager); err != nil {
		return nil, err
	}

	return &BuyWagerResponse{
		ID:          purchase.ID,
		WagerID:     purchase.WagerID,
		BuyingPrice: purchase.BuyingPrice,
		BoughtAt:    uint64(purchase.BoughtAt.Unix()),
	}, nil
}

func (s *serviceImpl) ListWagers(page uint32, limit uint32) (res []*WagerResponse, err error) {
	res = []*WagerResponse{}
	wagers, err := s.wagerRepository.FindAll(page, limit)
	if err != nil {
		return
	}

	for _, wager := range wagers {
		res = append(res, &WagerResponse{
			ID:                  wager.ID,
			TotalWagerValue:     wager.TotalWagerValue,
			Odds:                wager.Odds,
			SellingPercentage:   wager.SellingPercentage,
			SellingPrice:        wager.SellingPrice,
			CurrentSellingPrice: wager.CurrentSellingPrice,
			PercentageSold:      wager.PercentageSold,
			AmountSold:          wager.AmountSold,
			PlacedAt:            uint64(wager.PlacedAt.Unix()),
		})
	}

	return
}
