package server

import (
	"fmt"
	"net"

	"Auth/pkg/config"
	pb "Auth/pkg/pb/admin"
	pbemployer "Auth/pkg/pb/employer"
	pbjobseeker "Auth/pkg/pb/jobseeker"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, adminServer pb.AdminServer, employerServer pbemployer.EmployerServer, jobseekerServer pbjobseeker.JobSeekerServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminServer(newServer, adminServer)
	pbemployer.RegisterEmployerServer(newServer, employerServer)
	pbjobseeker.RegisterJobSeekerServer(newServer, jobseekerServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
