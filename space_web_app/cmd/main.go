package main

import (
	"os"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/grpc/client"
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

	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	newClient := pb.NewNasaClient(conn)

	gRPCClient := client.New(&newClient, logger)

	handlers := httphandlers.New(logger, gRPCClient)

	router := routers.New(handlers)

	router.SetUpHandlers()

	app := app.New(logger, cfg, router.Mux)

	app.HTTPServer.MustRun()
}
