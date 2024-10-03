// pkg/jwt/token.go

package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims map[string]interface{}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func GenerateToken(secretKey []byte, expirationTime time.Duration, claims TokenClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	standardClaims := token.Claims.(jwt.MapClaims)
	standardClaims["exp"] = time.Now().Add(expirationTime).Unix()

	for key, value := range claims {
		standardClaims[key] = value
	}

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
