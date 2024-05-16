package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/job"
	"fmt"

	"google.golang.org/grpc"
)

type chatClient struct {
	Client pb.JobClient
}

func NewChatClient(cfg config.Config) interfaces.JobClient {
	grpcConnection, err := grpc.Dial(cfg.HireoJob, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewJobClient(grpcConnection)

	return &jobClient{
		Client: grpcClient,
	}
}
