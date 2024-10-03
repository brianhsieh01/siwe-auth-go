package auth

import (
	"time"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/config"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/pkg/ethereum"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/pkg/jwt"
	"github.com/spruceid/siwe-go"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
}

func NewAuthService(repo *Repository, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}

type NonceResponse struct {
	Nonce string `json:"nonce"`
	Token string `json:"token"`
}

func (s *Service) GenerateNonce(address string) (*NonceResponse, error) {
	if err := ethereum.ValidateAddress(address); err != nil {
		return nil, err
	}

	nonce := siwe.GenerateNonce()

	claims := jwt.TokenClaims{
		"nonce":   nonce,
		"address": address,
	}

	tokenString, err := jwt.GenerateToken([]byte(s.cfg.JWTSecret), time.Minute*5, claims)
	if err != nil {
		return nil, err
	}

	return &NonceResponse{
		Nonce: nonce,
		Token: tokenString,
	}, nil
}
