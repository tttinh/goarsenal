package wager

import (
	"time"

	"github.com/tttinh/goarsenal/infra/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) CreateWager(req CreateWagerRequest) (res *WagerResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Infow("create_wager",
			"req", req,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = s.Service.CreateWager(req)

	return res, err
}

func (s *loggingService) BuyWager(wagerID uint32, req BuyWagerRequest) (res *BuyWagerResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Infow("buy_wager",
			"req", req,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = s.Service.BuyWager(wagerID, req)

	return
}

func (s *loggingService) ListWagers(page uint32, limit uint32) (res []*WagerResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Infow("list_wagers",
			"page", page,
			"limit", limit,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = s.Service.ListWagers(page, limit)

	return
}
