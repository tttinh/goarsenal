package wager

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/infra/log"
	"github.com/tttinh/goarsenal/repository"
)

type Service interface {
	CreateWager(req CreateWagerRequest) (*WagerResponse, error)
}

type Controller interface {
	CreateWager(c *gin.Context)
}

func NewService(repo repository.WagerRepository) Service {
	return &serviceImpl{repo}
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func NewController(service Service) *controllerImpl {
	return &controllerImpl{
		service: service,
	}
}
