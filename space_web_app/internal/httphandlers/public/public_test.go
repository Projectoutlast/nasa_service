package public

import (
	"log/slog"
	"net/http"
	"testing"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"github.com/stretchr/testify/assert"
)

type mockAuthClient struct {
	pb.AuthClient
}

func TestNew(t *testing.T) {
	authClient := &mockAuthClient{}

	publicHandlers := New(authClient, slog.Default())

	assert.IsType(t, &PublicHTTPHandlers{}, publicHandlers)
	assert.NotNil(t, publicHandlers.store)
	assert.Equal(t, publicHandlers.store.Options.Path, "/")
	assert.Equal(t, publicHandlers.store.Options.MaxAge, 3600*8)
	assert.Equal(t, publicHandlers.store.Options.HttpOnly, true)
	assert.Equal(t, publicHandlers.store.Options.Secure, false)
	assert.Equal(t, publicHandlers.store.Options.SameSite, http.SameSiteLaxMode)
}

func TestHandlerc(t *testing.T) {
	// authClient := &mockAuthClient{}

	// publicHandlers := New(authClient, slog.Default())
	
	// req, err := http.NewRequest("GET", "/", nil)
}
