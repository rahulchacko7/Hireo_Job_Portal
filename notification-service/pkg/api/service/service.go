package service

import (
	"context"
	"fmt"
	logging "notification/Logging"
	pb "notification/pkg/pb/noti"
	interfaces "notification/pkg/usecase/interface"
	"notification/pkg/utils/models"
	"os"

	"github.com/sirupsen/logrus"
)

type NotiServer struct {
	notiUsecase interfaces.NotiUseCase
	pb.UnimplementedNotificationServiceServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewnotiServer(usecase interfaces.NotiUseCase) pb.NotificationServiceServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	return &NotiServer{
		notiUsecase: usecase,
		Logger:      logger,
		LogFile:     logFile,
	}
}

func (ad *NotiServer) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	ad.Logger.Info("GetNotification at NotificationServer started")
	userid := req.UserID

	fmt.Println("at noti clinet servc", userid)

	result, err := ad.notiUsecase.GetNotification(int(userid), models.Pagination{Limit: int(req.Limit), Offset: int(req.Offset)})
	if err != nil {
		ad.Logger.Error("error from notificationUsecase", err)
		return nil, err
	}
	ad.Logger.Info("GetNotification at notificationUsecase success")
	var final []*pb.Message

	for _, v := range result {
		final = append(final, &pb.Message{
			UserId:   int64(v.UserID),
			Username: v.Username,
			Id:       int64(v.ID),
			Message:  v.Message,
			Time:     v.CreatedAt,
			PostId:   int64(v.PostID),
		})
	}
	ad.Logger.Info("GetNotification at NotificationServer success")
	return &pb.GetNotificationResponse{
		Notification: final,
	}, nil
}
