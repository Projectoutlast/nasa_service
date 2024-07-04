package httphandlers

import (
	"context"
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

var baseSpaceLayout = "./assets/html/base.layout.html"

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
}

func New(authClient AuthGRPCClient, log *slog.Logger, nasaClient NasaGRPCClient) *HTTPHandlers {
	return &HTTPHandlers{
		authClient: authClient,
		log:        log,
		nasaClient: nasaClient,
	}
}
