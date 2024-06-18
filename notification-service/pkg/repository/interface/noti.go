package interfaces

import "notification/pkg/utils/models"

type NotiRepository interface {
	StorenotificationReq(models.NotificationReq) error
	GetNotification(userid int, req models.Pagination) ([]models.Notification, error)
}
