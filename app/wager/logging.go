package wager

import (
	"time"

	"github.com/tttinh/goarsenal/infra/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) CreateGroup(requesterID string, req CreateWagerRequest) (err error) {
	defer func(begin time.Time) {
		s.logger.Infow("create_wager",
			"requester", requesterID,
			"req", req,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateWager(requesterID, req)
}
