package httphandlers

import (
	"log/slog"
	"net/http"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"github.com/gorilla/sessions"
)

var (
	baseSpaceLayout = "./assets/html/public/base.layout.html"
	baseUrl         = "http://localhost:50061"
)

type HTTPHandlers struct {
	authClient pb.AuthClient
	log        *slog.Logger
	store      *sessions.CookieStore
}

func New(authClient pb.AuthClient, log *slog.Logger) *HTTPHandlers {
	store := sessions.NewCookieStore([]byte("something-very-secret"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	return &HTTPHandlers{
		authClient: authClient,
		log:        log,
		store:      store,
	}
}
