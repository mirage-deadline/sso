package auth

import (
	"context"
	sso_v1 "github.com/mirage-deadline/sso-proto/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	sso_v1.UnimplementedAuthServer
}

func Register(grpc *grpc.Server) {
	sso_v1.RegisterAuthServer(grpc, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *sso_v1.LoginRequest) (*sso_v1.LoginResponse, error) {
	panic("unimplemented")
}

func (s *serverAPI) Register(ctx context.Context, req *sso_v1.RegisterRequest) (*sso_v1.RegisterResponse, error) {
	panic("unimplemented")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *sso_v1.IsAdminRequest) (*sso_v1.IsAdminResponse, error) {
	panic("unimplemented")
}
