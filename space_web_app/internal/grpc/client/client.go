package client

import (
	"context"
	"io"
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

type NasaGRPCClient struct {
	client pb.NasaClient
	log    *slog.Logger
}

func New(client *pb.NasaClient, log *slog.Logger) *NasaGRPCClient {
	return &NasaGRPCClient{
		*client,
		log,
	}
}

func (c *NasaGRPCClient) GetRandomSpaseImage(ctx context.Context, req *pb.RandomSpaseImageRequest) (*pb.RandomSpaseImageResponse, error) {
	stream, err := c.client.RandomSpaseImage(ctx, req)

	if err != nil {
		c.log.Error(err.Error())
	}

	response := &pb.RandomSpaseImageResponse{}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.log.Error(err.Error())

			return nil, err
		}

		response.Copyright = res.Copyright
		response.Date = res.Date
		response.Explanation = res.Explanation
		response.Title = res.Title
		response.Data = res.Data
	}

	return response, nil
}
