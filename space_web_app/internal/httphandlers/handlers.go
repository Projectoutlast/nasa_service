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

type HTTPHandlers struct {
	log        *slog.Logger
	nasaClient NasaGRPCClient
}

func New(log *slog.Logger, nasaClient NasaGRPCClient) *HTTPHandlers {
	return &HTTPHandlers{
		log:        log,
		nasaClient: nasaClient,
	}
}
