package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	authGrpc "github.com/Projectoutlast/space_service/auth_service/internal/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	log        *slog.Logger
	port       int
}

func New(log *slog.Logger, authService authGrpc.AuthUsecase, port int) *App {
	gRPCServer := grpc.NewServer()

	authGrpc.Register(gRPCServer, log, authService)

	return &App{
		gRPCServer: gRPCServer,
		log:        log,
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

	if err = a.gRPCServer.Serve(l); err != nil {
		a.log.Error(err.Error())
		return err
	}

	return nil
}
