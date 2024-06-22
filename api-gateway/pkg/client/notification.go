package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/noti"
	"HireoGateWay/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type notificationClient struct {
	client pb.NotificationServiceClient
}

func NewNotificationClient(cfg config.Config) interfaces.NotificationClient {
	grpcConnection, err := grpc.Dial(cfg.NotificationSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewNotificationServiceClient(grpcConnection)
	return &notificationClient{
		client: grpcClient,
	}
}

func (ad *notificationClient) GetNotification(userid int32, pagin models.NotificationPagination) ([]models.NotificationResponse, error) {
	data, err := ad.client.GetNotification(context.Background(), &pb.GetNotificationRequest{
		UserID: int64(userid),
		Limit:  int64(pagin.Limit),
		Offset: int64(pagin.Offset),
	})

	fmt.Println("at clinet ", userid)
	if err != nil {
		return []models.NotificationResponse{}, err

	}
	var response []models.NotificationResponse

	for _, v := range data.Notification {
		notificationresponse := models.NotificationResponse{
			UserID: int(v.UserId),
			// Username:  v.Username,
			// Profile:   v.Profile,
			Message:   v.Message,
			CreatedAt: v.Time,
		}
		response = append(response, notificationresponse)
	}
	return response, nil
}
