package main

import (
	"os"

	"github.com/Projectoutlast/space_service/auth_service/internal/app"
	"github.com/Projectoutlast/space_service/auth_service/internal/config"
	"github.com/Projectoutlast/space_service/auth_service/internal/logging"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	newLogger := logging.New(cfg.Environment, os.Stdout)

	application := app.New(newLogger, cfg)

	application.GRPCServer.MustRun()
}
