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
			"res", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = s.Service.CreateWager(req)

	return res, err
}
