package app

import (
	"log/slog"

	grpcapp "github.com/Projectoutlast/space_service/nasa_grpc_service/internal/app/grpc"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/config"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/services/nasa"

	nasaAPI "github.com/Projectoutlast/space_service/nasa_grpc_service/internal/third_parties_api/nasa"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	newNasaAPI := nasaAPI.New(log, cfg)

	nasaService := nasa.New(log, newNasaAPI)

	grpcApp := grpcapp.New(log, nasaService, cfg.GrpcConfig.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
