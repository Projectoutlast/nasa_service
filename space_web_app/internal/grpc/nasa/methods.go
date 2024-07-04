package nasa

import (
	"context"
	"io"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func (c *NasaGRPCClient) GetRandomSpaseImage(ctx context.Context, req *pb.RandomSpaseImageRequest) (*pb.RandomSpaseImageResponse, error) {
	stream, err := c.nasa.RandomSpaseImage(ctx, req)

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
