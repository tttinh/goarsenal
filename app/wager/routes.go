package wager

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, service Service) {
	ctrl := NewController(service)
	r.POST("/wagers", ctrl.CreateWager)
	r.POST("/buy/:wager_id", ctrl.BuyWager)
}
