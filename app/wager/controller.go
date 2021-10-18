package wager

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/infra/transport/http"
)

type controllerImpl struct {
	service Service
}

func (ctrl *controllerImpl) CreateWager(c *gin.Context) {
	var err error
	var req CreateWagerRequest

	err = http.BindAndValid(c, &req)
	if err != nil {
		http.BadRequest(c, err.Error())
		return
	}

	res, err := ctrl.service.CreateWager(req)

	if err != nil {
		http.BadRequest(c, err.Error())
	} else {
		http.Created(c, res)
	}
}

func (ctrl *controllerImpl) BuyWager(c *gin.Context) {
	var err error
	var req BuyWagerRequest

	err = http.BindAndValid(c, &req)
	if err != nil {
		http.BadRequest(c, err.Error())
	}

	wagerID, err := strconv.ParseUint(c.Param("wager_id"), 10, 32)
	if err != nil {
		http.BadRequest(c, "Bad wager id format!")
		return
	}

	res, err := ctrl.service.BuyWager(uint32(wagerID), req)
	if err != nil {
		http.BadRequest(c, err.Error())
	} else {
		http.Created(c, res)
	}
}
