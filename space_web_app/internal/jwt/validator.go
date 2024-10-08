package jwt

import (
	"crypto"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Validator struct {
	key crypto.PublicKey
}

func NewValidator(publicKeyPath string) (*Validator, error) {
	keyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key: %w", err)
	}

	key, err := jwt.ParseEdPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return &Validator{key: key}, nil
}

func (v *Validator) GetToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return v.key, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}
