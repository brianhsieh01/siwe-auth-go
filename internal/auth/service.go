package auth

import (
	"errors"
	"time"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/config"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/pkg/ethereum"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/pkg/jwt"
	"github.com/ethereum/go-ethereum/crypto"
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

type SignInResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Token     string `json:"token"`
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

func (s *Service) VerifySignature(req SignInRequest) (*SignInResponse, error) {
	siweMessage, err := siwe.ParseMessage(req.Message)
	if err != nil {
		return nil, errors.New("failed to parse SIWE message")
	}

	address := siweMessage.GetAddress().String()
	nonce := siweMessage.GetNonce()

	claims, err := jwt.VerifyToken(req.Token, []byte(s.cfg.JWTSecret))
	if err != nil {
		if errors.Is(err, jwt.ErrInvalidToken) {
			return nil, errors.New("invalid or expired token")
		}
		return nil, err
	}

	if claims["nonce"] != nonce || claims["address"] != address {
		return nil, errors.New("invalid nonce or address")
	}

	publicKey, err := siweMessage.Verify(req.Signature, nil, &nonce, nil)
	if err != nil {
		return nil, errors.New("failed to verify signature")
	}
	recoveredAddr := crypto.PubkeyToAddress(*publicKey)

	if address != recoveredAddr.Hex() {
		return nil, errors.New("recovered address does not match")
	}

	newClaims := jwt.TokenClaims{
		"address": address,
	}
	token, err := jwt.GenerateToken([]byte(s.cfg.JWTSecret), time.Hour*24, newClaims)
	if err != nil {
		return nil, errors.New("failed to generate session token")
	}

	return &SignInResponse{Token: token}, nil
}
