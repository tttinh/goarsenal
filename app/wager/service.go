package wager

import (
	"github.com/tttinh/goarsenal/entity"
	"github.com/tttinh/goarsenal/repository"
)

type serviceImpl struct {
	wagerRepository repository.WagerRepository
}

func (s *serviceImpl) CreateWager(requesterID string, req CreateWagerRequest) error {
	wager := &entity.Wager{
		TotalWagerValue:   req.TotalWagerValue,
		Odds:              req.Odds,
		SellingPercentage: req.SellingPercentage,
		SellingPrice:      req.SellingPrice,
	}

	return s.wagerRepository.AddWager(wager)
}
