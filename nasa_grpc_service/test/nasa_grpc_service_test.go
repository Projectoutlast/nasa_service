package test_test

import (
	"context"
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/app"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/config"
	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/logging"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func TestNasaGRPCService(t *testing.T) {
	// Set up nasa server
	go startTestNasaGRPCServer(t)
	client := newTestNasaGRPCClient(t, "localhost:50052")

	req := &pb.RandomSpaseImageRequest{}

	stream, err := client.RandomSpaseImage(context.Background(), req)
	require.NoError(t, err)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		_, file, line, _ := runtime.Caller(0)
		t.Logf("Error at %s:%d - %v", file, line, err)

		require.NoError(t, err)
		require.NotNil(t, res)
	}

}

func startTestNasaGRPCServer(_ *testing.T) {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Environment, os.Stdout)
	application := app.New(logger, cfg)
	application.GRPCServer.MustRun()
}

func newTestNasaGRPCClient(t *testing.T, serverAddress string) pb.NasaClient {
	conn, err := grpc.NewClient(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewNasaClient(conn)
}
