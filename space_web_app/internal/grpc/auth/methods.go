package auth

import (
	"context"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func (a *AuthGRPCClient) Register(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	res, err := a.auth.Registration(context.Background(), req)

	if err != nil {
		a.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (a *AuthGRPCClient) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {

	return nil, nil
}
