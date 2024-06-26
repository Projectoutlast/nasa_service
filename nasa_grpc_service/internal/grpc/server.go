package grpc

import (
	pb "github.com/Projectoutlast/nasa_proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NasaUsecase interface {
	GetRandomSpaseImage() (*pb.RandomSpaseImageResponse, error)
}

type serverAPI struct {
	pb.UnimplementedNasaServer
	usecase NasaUsecase
}

func Register(gRPCServer *grpc.Server, usecase NasaUsecase) {
	pb.RegisterNasaServer(gRPCServer, &serverAPI{usecase: usecase})
}

func (s *serverAPI) GetRandomSpaseImage(req *pb.RandomSpaseImageRequest, stream pb.Nasa_RandomSpaseImageServer) error {
	response, err := s.usecase.GetRandomSpaseImage()
	if err != nil {
		return status.Error(codes.Internal, codes.Internal.String())
	}

	if err := stream.Send(response); err != nil {
		return status.Error(codes.Internal, codes.Internal.String())
	}

	return nil
}
