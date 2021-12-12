package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type ServerOptions func(o *Server)

func NewServer(opts...ServerOptions) *Server {
	srv := &Server{
		ctx: context.Background(),
		network: "tcp",
		address: ":8080",
	}
	for _,o := range opts {
		o(srv)
	}
	srv.Server = grpc.NewServer()
	return srv
}

func Network(network string) ServerOptions {
	return func(o *Server) {
		o.network = network
	}
}

func Address(addr string) ServerOptions {
	return func(o *Server) {
		o.address = addr
	}
}

type Server struct {
	*grpc.Server
	ctx			context.Context
	lis 		net.Listener
	once 		sync.Once
	err 		error
	network 	string
	address 	string
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.Endpoint(); err != nil {
		return err
	}
	return s.Serve(s.lis)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.lis.Close()
}

func (s *Server) Endpoint() error {
	s.once.Do(func() {
		lis,err := net.Listen(s.network,s.address)
		if err != nil {
			s.err = err
			return
		}
		s.lis = lis
	})
	return s.err
}

func (s *Server) GetContext() context.Context {
	return s.ctx
}