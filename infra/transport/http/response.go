package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string
}

// BadRequest returns bad request error.
func BadRequest(c *gin.Context, errMessage string) {
	c.JSON(http.StatusBadRequest, errorResponse{Error: errMessage})
}

// Ok returns success.
func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Created returns success.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}
