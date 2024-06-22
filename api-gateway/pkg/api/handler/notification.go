package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	GRPC_Client interfaces.NotificationClient
}

func NewNotificationHandler(notiClient interfaces.NotificationClient) *NotificationHandler {
	return &NotificationHandler{
		GRPC_Client: notiClient,
	}
}

func (ad *NotificationHandler) GetNotification(c *gin.Context) {
	var notificationRequest models.NotificationPagination
	if err := c.ShouldBindJSON(&notificationRequest); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "details given in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	// Retrieve employerID from context
	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Debug: Print employerID type and value
	fmt.Printf("employerID type: %T, value: %v\n", employerID, employerID)

	// Attempt to convert employerID to string
	// id, ok := employerID.(int)
	// if !ok {
	// 	errs := response.ClientResponse(http.StatusBadRequest, "Failed to convert employer ID to string", nil, nil)
	// 	c.JSON(http.StatusBadRequest, errs)
	// 	return
	// }

	// id, err := strconv.Atoi(newID)
	// if err != nil {
	// 	errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errs)
	// 	return
	// }

	// fmt.Println("new id", id)

	// fmt.Println("id handler after conversion", id)

	// Call ad.GRPC_Client.GetNotification with id and notificationRequest
	result, errs := ad.GRPC_Client.GetNotification(employerID.(int32), notificationRequest)
	if errs != nil {
		errss := response.ClientResponse(http.StatusBadRequest, "Error in getting notification", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errss)
		return
	}

	// Return success response
	succesres := response.ClientResponse(http.StatusOK, "Successfully got all notifications", result, nil)
	c.JSON(http.StatusOK, succesres)
}
