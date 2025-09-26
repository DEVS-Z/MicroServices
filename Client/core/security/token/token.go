package token

import (
	"errors"
	"main/core/security/key"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ValidToken(r string) (*jwt.Token, error) {
	tokenStr := getToken(r)
	if tokenStr == "" {
		return nil, errors.New("missing or malformed token")
	}

	publicKey, err := key.LoadPublicKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func getToken(token string) string {
	parts := strings.Split(token, " ")
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		cleaned := strings.Trim(parts[1], "\"")
		return cleaned
	}
	return ""
}
