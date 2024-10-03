package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewAuthHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type ProfileResponse struct {
	Address string `json:"address"`
}

func (h *Handler) GetNonce(c echo.Context) error {
	address := c.QueryParam("address")
	response, err := h.service.GenerateNonce(address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}
