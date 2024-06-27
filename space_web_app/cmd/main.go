package main

import (
	"os"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)

	router := mux.NewRouter()

	app := app.New(logger, cfg, router)

	app.HTTPServer.MustRun()
}
