package server

import (
	"fmt"
	"net"

	"Auth/pkg/config"
	pb "Auth/pkg/pb/admin"
	pbemployer "Auth/pkg/pb/employer"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, adminServer pb.AdminServer, employerServer pbemployer.EmployerServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminServer(newServer, adminServer)
	pbemployer.RegisterEmployerServer(newServer, employerServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
