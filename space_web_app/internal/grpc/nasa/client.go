package nasa

import (
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

type NasaGRPCClient struct {
	nasa pb.NasaClient
	log  *slog.Logger
}

func New(client *pb.NasaClient, log *slog.Logger) *NasaGRPCClient {
	return &NasaGRPCClient{
		*client,
		log,
	}
}
