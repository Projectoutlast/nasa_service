package httphandlers

import (
	"context"
	"log/slog"
	"net/http"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"github.com/gorilla/sessions"
)

var (
	baseSpaceLayout = "./assets/html/base.layout.html"
	baseUrl         = "http://localhost:50061"
)

type NasaGRPCClient interface {
	GetRandomSpaseImage(context.Context, *pb.RandomSpaseImageRequest) (*pb.RandomSpaseImageResponse, error)
}

type AuthGRPCClient interface {
	Register(context.Context, *pb.RegistrationRequest) (*pb.RegistrationResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
}

type HTTPHandlers struct {
	authClient AuthGRPCClient
	log        *slog.Logger
	nasaClient NasaGRPCClient
	store      *sessions.CookieStore
}

func New(authClient AuthGRPCClient, log *slog.Logger, nasaClient NasaGRPCClient) *HTTPHandlers {
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
		nasaClient: nasaClient,
		store:      store,
	}
}
