package server

import (
	v1 "github.com/kratos-portal/api/user/v1"
	"github.com/kratos-portal/internal/user/service"
	"github.com/kratos-portal/pkg/transport/grpc"
)

// NewGRPCServer new a gRPC server
func NewGRPCServer(network,addr string,user *service.UserService) *grpc.Server {
	var opts []grpc.ServerOptions
	if network != "" {
		opts = append(opts,grpc.Network(network))
	}
	if addr != "" {
		opts = append(opts,grpc.Address(addr))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServiceServer(srv,user)
	return srv
}
