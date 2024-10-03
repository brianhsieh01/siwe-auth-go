package auth

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

func (h *Handler) SignIn(c echo.Context) error {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	response, err := h.service.VerifySignature(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetProfile(c echo.Context) error {
	userClaims, ok := c.Get("jwtClaims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	formattedClaims, err := json.MarshalIndent(userClaims, "", "  ")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to format claims"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile retrieved successfully",
		"claims":  json.RawMessage(formattedClaims),
	})
}
