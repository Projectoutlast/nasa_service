package app

import (
	"log/slog"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app/httpapp"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/gorilla/mux"
)

type App struct {
	HTTPServer *httpapp.App
}

func New(log *slog.Logger, cfg *config.Config, router *mux.Router) *App {

	httpApp := httpapp.New(log, cfg, router) 
	return &App{
		HTTPServer: httpApp,
	}
}
