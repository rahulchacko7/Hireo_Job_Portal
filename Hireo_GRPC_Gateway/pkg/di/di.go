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

	serverHTTP := server.NewServerHTTP(adminHandler)

	return serverHTTP, nil
}
