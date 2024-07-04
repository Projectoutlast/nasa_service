package main

import (
	"os"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/grpc/auth"
	"github.com/Projectoutlast/space_service/space_web_app/internal/grpc/nasa"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
	"github.com/Projectoutlast/space_service/space_web_app/internal/routers"
	"google.golang.org/grpc"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)

	connNasa, err := grpc.NewClient(cfg.ClientsAddress.Nasa, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connNasa.Close()
	nasaClient := pb.NewNasaClient(connNasa)

	connAuth, err := grpc.NewClient(cfg.ClientsAddress.Auth, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connAuth.Close()
	authClient := pb.NewAuthClient(connAuth)

	nasaGRPCClient := nasa.New(&nasaClient, logger)
	authGRPCClient := auth.New(&authClient, logger)

	handlers := httphandlers.New(authGRPCClient, logger, nasaGRPCClient)

	router := routers.New(handlers, cfg.Server.FileServerDir, cfg.Server.StaticPrefix)

	router.SetUpHandlers()
	router.SetUpFileServer()

	app := app.New(logger, cfg, router.Mux)

	app.HTTPServer.MustRun()
}
