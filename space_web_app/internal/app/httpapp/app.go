package httpapp

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
)

type App struct {
	config     *config.Config
	httpServer *http.Server
	log        *slog.Logger
	router     *mux.Router
}

func New(log *slog.Logger, config *config.Config, router *mux.Router) *App {
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      router,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
	}

	return &App{
		config:     config,
		httpServer: httpServer,
		log:        log,
		router:     router,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	a.log.Info(fmt.Sprintf("HTTP server is running on port %s", a.httpServer.Addr))

	if err := a.httpServer.ListenAndServeTLS(a.config.CertFile, a.config.KeyFile); err != nil {
		a.log.Error(err.Error())
		return err
	}

	return nil
}
