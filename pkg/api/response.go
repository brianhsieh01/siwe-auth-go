package api

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewSuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(c echo.Context, statusCode int, message string, err error) error {
	return c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Error:   err.Error(),
	})
}
