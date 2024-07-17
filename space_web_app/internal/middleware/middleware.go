package middleware

import (
	"log/slog"

	"github.com/Projectoutlast/space_service/space_web_app/internal/jwt"
)

type Middleware struct {
	log       *slog.Logger
	validator *jwt.Validator
}

func New(log *slog.Logger, validator *jwt.Validator) *Middleware {
	return &Middleware{
		log:       log,
		validator: validator,
	}
}
