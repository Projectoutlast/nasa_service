package test_test

import (
	"log/slog"
	"testing"

	jwtIssuer "github.com/Projectoutlast/space_service/auth_service/internal/services/jwt"
	"github.com/stretchr/testify/require"
)

func TestUsecase(t *testing.T) {
	_, err := jwtIssuer.NewIssuer("hello", slog.Default())
	require.Error(t, err)

}
