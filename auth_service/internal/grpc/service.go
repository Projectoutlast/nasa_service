package grpc

import (
	"context"
	"log/slog"

	pb "github.com/Projectoutlast/nasa_proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthUsecase interface {
	Registration(string, string) (int64, error)
	Login(string, string) (string, error)
}

type AuthService struct {
	pb.UnimplementedAuthServer

	log     *slog.Logger
	usecase AuthUsecase
}

func Register(gRPCServer *grpc.Server, log *slog.Logger, usecase AuthUsecase) {
	pb.RegisterAuthServer(gRPCServer, &AuthService{log: log, usecase: usecase})
}

func (a *AuthService) Registration(ctx context.Context, req *pb.RegistrationRequest) (res *pb.RegistrationResponse, err error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		a.log.Error("email or password is empty")
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	user_id, err := a.usecase.Registration(req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return &pb.RegistrationResponse{UserId: user_id}, nil
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (res *pb.LoginResponse, err error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		a.log.Error("email or password is empty")
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	token, err := a.usecase.Login(req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return &pb.LoginResponse{Token: token}, nil
}
