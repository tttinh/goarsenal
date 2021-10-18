package http

import (
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, req interface{}) error {
	return c.Bind(req)
}
