package interfaces

import "notification/pkg/utils/models"

type NotiUseCase interface {
	ConsumeNotification()
	GetNotification(userid int, mod models.Pagination) ([]models.NotificationResponse, error)
}
