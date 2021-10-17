package wager

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, service Service) {
	router := r.Group("/api/v1/wagers")

	ctrl := NewController(service)
	router.POST("/", ctrl.CreateWager)
}
