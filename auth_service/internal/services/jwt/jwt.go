package services

import (
	"crypto"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Issuer struct {
	key crypto.PrivateKey
	log *slog.Logger
}

func NewIssuer(privateKeyPath string, log *slog.Logger) (*Issuer, error) {
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	key, err := jwt.ParseEdPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &Issuer{key: key, log: log}, nil
}

func (i *Issuer) IssueToken(email string, services []string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, jwt.MapClaims{
		"exp":   now.Add(time.Hour * 24).Unix(),
		"email": email,
		"roles": services,
	})

	tokenString, err := token.SignedString(i.key)
	if err != nil {
		i.log.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}
