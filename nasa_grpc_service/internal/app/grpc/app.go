package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	nasagrpc "github.com/Projectoutlast/space_service/nasa_grpc_service/internal/grpc"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/interceptor"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, nasaService nasagrpc.NasaUsecase, port int) *App {
	gRPCServer := grpc.NewServer(
		grpc.StreamInterceptor(interceptor.NewStreamInterceptor(log)),
	)

	nasagrpc.Register(gRPCServer, nasaService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.Error(err.Error())
		return err
	}

	a.log.Info(fmt.Sprintf("gRPC server is running on port %d", a.port))

	if err = a.gRPCServer.Serve(l);	err != nil {
		a.log.Error(err.Error())
		return err
	}

	return nil
}
