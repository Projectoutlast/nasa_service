package middleware

import (
	"log/slog"

	"github.com/Projectoutlast/space_service/space_web_app/internal/jwt"
	"github.com/gorilla/sessions"
)

type Middleware struct {
	log       *slog.Logger
	store     *sessions.CookieStore
	validator *jwt.Validator
}

func New(log *slog.Logger, store *sessions.CookieStore, validator *jwt.Validator) *Middleware {
	return &Middleware{
		log:       log,
		store:     store,
		validator: validator,
	}
}
