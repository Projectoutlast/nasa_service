package test_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/Projectoutlast/space_service/space_web_app/internal/app"
	"github.com/Projectoutlast/space_service/space_web_app/internal/config"
	"github.com/Projectoutlast/space_service/space_web_app/internal/logging"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestSuccessStartSpaceWebServer(t *testing.T) {
	t.Parallel()

	logger, cfg, router := preparationStartSpaceWebServer(t)

	successStartApplication := app.New(logger, cfg, router)

	var err error

	go func() {
		err = successStartApplication.HTTPServer.Run()
	}()

	require.Nil(t, err)
}

func TestCreateLogger(t *testing.T) {
	t.Parallel()

	logger := logging.New("testing", os.Stdout)
	require.NotNil(t, logger)

	logger = logging.New("production", os.Stdout)
	require.NotNil(t, logger)
}

func preparationStartSpaceWebServer(_ *testing.T) (*slog.Logger, *config.Config, *mux.Router) {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)

	r := mux.NewRouter()

	return logger, cfg, r
}
