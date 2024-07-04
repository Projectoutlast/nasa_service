package test_test

import (
	"log/slog"
	"testing"

	"github.com/Projectoutlast/space_service/auth_service/internal/config"
	"github.com/stretchr/testify/require"

	"github.com/Projectoutlast/space_service/auth_service/internal/app"
)

func TestAppPanicAndError(t *testing.T) {
	cfg := config.Config{
		PKeyPath: "test",
		AuthConfig: config.Auth{
			Port: 0,
		},
	}

	require.Panics(
		t,
		func() { app.New(slog.Default(), &cfg, nil) },
	)

}
