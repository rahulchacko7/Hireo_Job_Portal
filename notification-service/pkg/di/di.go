package di

import (
	server "notification/pkg/api"
	"notification/pkg/api/service"
	"notification/pkg/client"
	"notification/pkg/config"
	"notification/pkg/db"
	"notification/pkg/repository"
	"notification/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	notiRepository := repository.NewnotiRepository(gormDB)
	noticlient := client.NewAuthClient(&cfg)
	noriUseCase := usecase.NewnotiUsecase(notiRepository, noticlient)
	notiServiceServer := service.NewnotiServer(noriUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notiServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go noriUseCase.ConsumeNotification()
	return grpcserver, nil
}
