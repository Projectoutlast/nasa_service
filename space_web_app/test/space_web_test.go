package test_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/public"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers/secure"
	"github.com/Projectoutlast/space_service/space_web_app/internal/jwt"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
	"github.com/Projectoutlast/space_service/space_web_app/internal/middleware"
	"github.com/Projectoutlast/space_service/space_web_app/internal/routers"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func TestMain(t *testing.T) {
	t.Parallel()

	// Set up before start server
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

	_ = pb.NewNasaClient(connNasa)
	authCliet := pb.NewAuthClient(connAuth)

	publicHandlers := public.New(authCliet, logger)
	secureHandlers := secure.New(logger, nasaClient)

	validator, err := jwt.NewValidator(cfg.PubKeyPath)
	if err != nil {
		panic(err)
	}

	newMiddleware := middleware.New(logger, validator)

	r := routers.New(publicHandlers, secureHandlers, cfg.Server.FileServerDir, newMiddleware, cfg.Server.StaticPrefix)
	r.SetUpHandlers()

	// Error in start server
	cfg.Server.Port = -4
	errStartServer := app.New(logger, cfg, r.Mux)
	err = errStartServer.HTTPServer.Run()
	require.Error(t, err)

	// Panic test
	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("panic expected")
			}
		}()
		errStartServer.HTTPServer.MustRun()
	})

	// Success start server
	cfg.Server.Port = 50061

	application := app.New(logger, cfg, r.Mux)
	go application.HTTPServer.MustRun()

	res, err := http.Get("http://localhost:50061" + "/random")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateLogger(t *testing.T) {
	t.Parallel()

	logger := logging.New("testing", os.Stdout)
	require.NotNil(t, logger)

	logger = logging.New("production", os.Stdout)
	require.NotNil(t, logger)
}
