package app

import (
	"log/slog"

	grpcapp "github.com/Projectoutlast/space_service/auth_service/internal/app/grpc"
	"github.com/Projectoutlast/space_service/auth_service/internal/config"
	"github.com/Projectoutlast/space_service/auth_service/internal/services"
	jwtIssuer "github.com/Projectoutlast/space_service/auth_service/internal/services/jwt"
	storage "github.com/Projectoutlast/space_service/auth_service/internal/storage/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {

	issuer, err := jwtIssuer.NewIssuer(cfg.PKeyPath, log)
	if err != nil {
		panic(err)
	}

	newStorage, err := storage.New(log, cfg.StoragePath)
	if err != nil {
		panic(err)
	}

	authServices := services.New(issuer, log, newStorage)

	grpcApp := grpcapp.New(log, authServices, cfg.AuthConfig.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
