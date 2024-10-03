package auth

import "errors"

type NonceResponse struct {
	Nonce string `json:"nonce"`
	Token string `json:"token"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Token     string `json:"token"`
}

// Custom errors
var (
	ErrGenerateNonce        = errors.New("failed to generate nonce")
	ErrInvalidAddressFormat = errors.New("invalid address format")
)
