package services

import (
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	"github.com/Projectoutlast/space_service/auth_service/internal/models/storage"
	jwtIssuer "github.com/Projectoutlast/space_service/auth_service/internal/services/jwt"
)

type SQLiteStorage interface {
	Registration(string, string) (int64, error)
	GetUser(string) (*storage.User, error)
	GetUserServices(string) ([]string, error)
}

type AuthUsecase struct {
	issuer  *jwtIssuer.Issuer
	log     *slog.Logger
	storage SQLiteStorage
}

func New(issuer *jwtIssuer.Issuer, log *slog.Logger, storage SQLiteStorage) *AuthUsecase {
	return &AuthUsecase{
		issuer:  issuer,
		log:     log,
		storage: storage,
	}
}

func (a *AuthUsecase) Registration(email, password string) (int64, error) {

	hashedPassword, err := a.hashPassword(password)

	if err != nil {
		return 0, err
	}

	user_id, err := a.storage.Registration(email, hashedPassword)

	if err != nil {
		return 0, err
	}

	return user_id, err
}

func (a *AuthUsecase) Login(email, password string) (string, error) {
	user, err := a.storage.GetUser(email)
	if err != nil {
		return "", err
	}

	if !a.checkPasswordHash(password, user.Hash) {
		return "", errors.New("invalid password")
	}

	services, err := a.storage.GetUserServices(user.Email)
	if err != nil {
		return "", err
	}

	token, err := a.issuer.IssueToken(user.Email, services)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}

	return token, nil
}

func (a *AuthUsecase) hashPassword(password string) (string, error) {
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}

	return string(bytesPassword), nil
}

func (a *AuthUsecase) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
