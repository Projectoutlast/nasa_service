package main

import (
	"os"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/authenticator"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/public"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/secure"
	"github.com/Projectoutlast/space_service/space_web_app/internal/interceptors"
	"github.com/Projectoutlast/space_service/space_web_app/internal/jwt"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
	"github.com/Projectoutlast/space_service/space_web_app/internal/middleware"
	"github.com/Projectoutlast/space_service/space_web_app/internal/routers"
	httpsessions "github.com/Projectoutlast/space_service/space_web_app/internal/http_sessions"
	"google.golang.org/grpc"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)

	connNasa, err := grpc.NewClient(cfg.ClientsAddress.Nasa, grpc.WithInsecure(), grpc.WithStreamInterceptor(interceptors.StreamLoggingInterceptor(logger)))
	if err != nil {
		panic(err)
	}
	defer connNasa.Close()
	nasaClient := pb.NewNasaClient(connNasa)

	connAuth, err := grpc.NewClient(cfg.ClientsAddress.Auth, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptors.UnaryLoggingInterceptor(logger)))
	if err != nil {
		panic(err)
	}
	defer connAuth.Close()
	authClient := pb.NewAuthClient(connAuth)

	httpSessions := httpsessions.New()

	publicHandlers := public.New(authClient, logger, httpSessions)
	secureHandlers := secure.New(logger, nasaClient)

	validator, err := jwt.NewValidator(cfg.PubKeyPath)
	if err != nil {
		panic(err)
	}

	newMiddleware := middleware.New(logger, httpSessions, validator)

	auth, err := authenticator.New()
	if err != nil {
		panic(err)
	}

	router := routers.New(auth, publicHandlers, secureHandlers, cfg.Server.FileServerDir, newMiddleware, cfg.Server.StaticPrefix)

	router.SetUpHandlers()
	router.SetUpFileServer()

	app := app.New(logger, cfg, router.Mux)

	app.HTTPServer.MustRun()
}
