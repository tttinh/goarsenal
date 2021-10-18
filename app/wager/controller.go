package wager

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/infra/transport/http"
)

type controllerImpl struct {
	service Service
}

func (ctrl *controllerImpl) ListWagers(c *gin.Context) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "0"), 10, 32)
	if err != nil {
		http.BadRequest(c, "Invalid `page` parameter: "+err.Error())
		return
	}

	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 32)
	if err != nil {
		http.BadRequest(c, "Invalid `limit` parameter: "+err.Error())
		return
	}

	res, err := ctrl.service.ListWagers(uint32(page), uint32(limit))

	if err != nil {
		http.BadRequest(c, err.Error())
	} else {
		http.Ok(c, res)
	}
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
		return
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
