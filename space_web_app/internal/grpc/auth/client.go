package auth

import (
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

type AuthGRPCClient struct {
	auth pb.AuthClient
	log  *slog.Logger
}

func New(client *pb.AuthClient, log *slog.Logger) *AuthGRPCClient {
	return &AuthGRPCClient{
		*client,
		log,
	}
}
