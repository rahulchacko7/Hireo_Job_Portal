package interfaces

import "HireoGateWay/pkg/utils/models"

type NotificationClient interface {
	GetNotification(userid int32, req models.NotificationPagination) ([]models.NotificationResponse, error)
}
