package test_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/grpc/client"
	"github.com/Projectoutlast/space_service/space_web_app/internal/httphandlers"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
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
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	newClient := pb.NewNasaClient(conn)
	gRPCClient := client.New(&newClient, logger)
	handlers := httphandlers.New(logger, gRPCClient)

	r := routers.New(handlers, cfg.Server.FileServerDir, cfg.Server.StaticPrefix)
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
