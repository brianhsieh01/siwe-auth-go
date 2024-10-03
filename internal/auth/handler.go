package auth

import (
	"net/http"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/pkg/api"
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
		switch err {
		case ErrInvalidAddressFormat:
			return api.NewErrorResponse(c, http.StatusBadRequest, "Invalid address", err)
		case ErrGenerateNonce:
			return api.NewErrorResponse(c, http.StatusInternalServerError, "Failed to generate nonce", err)
		default:
			return api.NewErrorResponse(c, http.StatusInternalServerError, "Unexpected error", err)
		}
	}
	return api.NewSuccessResponse(c, http.StatusOK, "Nonce generated successfully", response)
}

func (h *Handler) SignIn(c echo.Context) error {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		return api.NewErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	response, err := h.service.VerifySignature(req)
	if err != nil {
		return api.NewErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err)
	}

	return api.NewSuccessResponse(c, http.StatusOK, "Signed in successfully", response)
}

func (h *Handler) GetProfile(c echo.Context) error {
	userClaims, ok := c.Get("jwtClaims").(jwt.MapClaims)
	if !ok {
		return api.NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	return api.NewSuccessResponse(c, http.StatusOK, "Profile retrieved successfully", userClaims)
}
