package authgrpc

import (
	"context"

	ssov1 "github.com/Underwoodoff/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAythServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAythServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
