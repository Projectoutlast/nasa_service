package secure

import (
	"log/slog"
	"net/http"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"github.com/gorilla/sessions"
)

var (
	baseSpaceLayout = "./assets/html/secure/base.layout.html"
	baseUrl         = "http://localhost:50061"
)

type SecureHTTPHandlers struct {
	log        *slog.Logger
	nasaClient pb.NasaClient
	store      *sessions.CookieStore
}

func New(log *slog.Logger, nasaClient pb.NasaClient) *SecureHTTPHandlers {
	store := sessions.NewCookieStore([]byte("something-very-secret"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	return &SecureHTTPHandlers{
		log:        log,
		nasaClient: nasaClient,
		store:      store,
	}
}
