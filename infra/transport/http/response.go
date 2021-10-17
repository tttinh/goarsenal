package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response400 returns bad request error.
func Response400(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, data)
}

// Response200 returns success.
func Response200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Response201 returns success.
func Response201(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}
