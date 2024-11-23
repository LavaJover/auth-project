package authgrpc

import (
	"context"
	"google.golang.org/grpc"
	"github.com/LavaJover/auth-project/auth-service/gen/authpb"
)

type serverAPI struct{
	authpb.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server){
	authpb.RegisterAuthServer(gRPC, serverAPI{})
}

func (s serverAPI) Login(
	 ctx context.Context,
	 loginRequest *authpb.LoginRequest,
) (*authpb.LoginResponse, error){
	panic("Implement me!")
}

func (s serverAPI) Register(
	ctx context.Context,
	registerRequest *authpb.RegisterRequest,
)(*authpb.RegisterResponse, error){
	panic("Implement me!")
}

func (s serverAPI) IsAdmin(
	ctx context.Context,
	isAdminRequest *authpb.IsAdminRequest,
) (*authpb.IsAdminResponse, error){
	panic("Implement me")
}