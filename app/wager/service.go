package wager

import (
	"github.com/tttinh/goarsenal/entity"
	"github.com/tttinh/goarsenal/repository"
)

type serviceImpl struct {
	wagerRepository    repository.WagerRepository
	purchaseRepository repository.PurchaseRepository
}

func (s *serviceImpl) CreateWager(req CreateWagerRequest) (*WagerResponse, error) {
	wager := &entity.Wager{
		TotalWagerValue:     req.TotalWagerValue,
		Odds:                req.Odds,
		SellingPercentage:   req.SellingPercentage,
		SellingPrice:        req.SellingPrice,
		CurrentSellingPrice: req.SellingPrice,
	}

	if err := s.wagerRepository.Save(wager); err != nil {
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

	purchase := &entity.Purchase{
		BuyingPrice: req.BuyingPrice,
		WagerID:     wagerID,
		Wager:       wager,
	}

	if err := s.purchaseRepository.Save(purchase); err != nil {
		return nil, err
	}

	return &BuyWagerResponse{
		ID: purchase.ID,
	}, nil
}
