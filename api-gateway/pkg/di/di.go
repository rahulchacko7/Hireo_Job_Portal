package di

import (
	server "HireoGateWay/pkg/api"
	"HireoGateWay/pkg/api/handler"
	"HireoGateWay/pkg/client"
	"HireoGateWay/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	employerClient := client.NewEmployerClient(cfg)
	employerHandler := handler.NewEmployerHandler(employerClient)

	jobSeekerClient := client.NewJobSeekerClient(cfg)
	jobSeekerHandler := handler.NewJobSeekerHandler(jobSeekerClient)

	serverHTTP := server.NewServerHTTP(adminHandler, employerHandler, jobSeekerHandler)

	return serverHTTP, nil
}
