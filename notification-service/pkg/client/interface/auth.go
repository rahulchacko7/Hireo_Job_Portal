package interfaces

import "notification/pkg/utils/models"

type Newauthclient interface {
	UserData(userid int) (models.UserData, error)
}
