package public

import (
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"github.com/gorilla/sessions"
)

var (
	baseSpaceLayout = "./assets/html/public/base.layout.html"
	baseUrl         = "http://localhost:50061"
)

type PublicHTTPHandlers struct {
	authClient pb.AuthClient
	log        *slog.Logger
	store      *sessions.CookieStore
}

func New(authClient pb.AuthClient, log *slog.Logger, store *sessions.CookieStore) *PublicHTTPHandlers {
	return &PublicHTTPHandlers{
		authClient: authClient,
		log:        log,
		store:      store,
	}
}
