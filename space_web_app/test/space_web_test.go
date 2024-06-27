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

func TestMain(t *testing.T) {
	_ = logging.New("testing", os.Stdout)
	_ = logging.New("production", os.Stdout)

	// Preparation before test
	logger, cfg, router := preparationStartSpaceWebServer(t)

	// Testing error start http server
	cfg.Server.Port = -4
	errorStartApplication := app.New(logger, cfg, router)
	require.Error(t, errorStartApplication.HTTPServer.Run())

	// // Testing panic in start http server
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic")
		}
	}()

	errorStartApplication.HTTPServer.MustRun()

	// Testing normal start http server
	cfg.Server.Port = 50061
	successStartApplication := app.New(logger, cfg, router)
	go successStartApplication.HTTPServer.MustRun()

	var err error
	go func() {
		err = successStartApplication.HTTPServer.Run()
	}()

	require.Nil(t, err)

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
