package wager

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/infra/log"
	"github.com/tttinh/goarsenal/repository"
)

type Service interface {
	CreateWager(req CreateWagerRequest) (*WagerResponse, error)
	BuyWager(wagerID uint32, req BuyWagerRequest) (*BuyWagerResponse, error)
	ListWagers(page uint32, limit uint32) ([]*WagerResponse, error)
}

type Controller interface {
	CreateWager(c *gin.Context)
	BuyWager(c *gin.Context)
	ListWagers(c *gin.Context)
}

func NewService(wagerRepo repository.WagerRepository, purchaseRepo repository.PurchaseRepository) Service {
	return &serviceImpl{wagerRepository: wagerRepo, purchaseRepository: purchaseRepo}
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func NewController(service Service) *controllerImpl {
	return &controllerImpl{
		service: service,
	}
}
