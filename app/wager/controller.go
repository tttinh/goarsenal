package wager

import (
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
	}

	res, err := ctrl.service.CreateWager(req)

	if err != nil {
		http.BadRequest(c, err.Error())
	} else {
		http.Created(c, res)
	}
}
