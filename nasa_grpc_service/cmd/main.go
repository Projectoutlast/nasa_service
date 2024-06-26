package main

import (
	"os"

	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/app"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/config"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/logging"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)

	application := app.New(logger, cfg)

	application.GRPCServer.MustRun()
}
