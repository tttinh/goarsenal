package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/infra/errorcode"
)

type errorResponse struct {
	Error errorcode.Code `json:"error"`
}

// ServerError returns internal server error.
func ServerError(c *gin.Context, ec errorcode.Code) {
	c.JSON(http.StatusInternalServerError, errorResponse{Error: ec})
}

// BadRequest returns bad request error.
func BadRequest(c *gin.Context, ec errorcode.Code) {
	c.JSON(http.StatusBadRequest, errorResponse{Error: ec})
}

// Ok returns success.
func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Created returns success.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}
