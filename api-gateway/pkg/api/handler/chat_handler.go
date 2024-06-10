package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
	helper      *helper.Helper
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
	}
}

// WebSocket
func (ch *ChatHandler) EmployerMessage(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "EmployerMessage")
	logEntry.Info("Processing EmployerMessage request")

	tokenString := c.Request.Header.Get("Authorization")
	logEntry.Info("Extracted Authorization header")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		logEntry.Error("Missing Authorization header")
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])
	logEntry.WithFields(logrus.Fields{
		"userID": userID,
		"error":  err,
	}).Info("Validated token")

	if err != nil {
		logEntry.WithError(err).Error("Invalid token")
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	logEntry.Info("Upgrading to WebSocket connection")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logEntry.WithError(err).Error("WebSocket Connection Issue")
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		delete(User, strconv.Itoa(int(userID.Id)))
		conn.Close()
		logEntry.Info("WebSocket connection closed and user removed from map")
	}()

	user := strconv.Itoa(int(userID.Id))
	User[user] = conn
	logEntry.WithField("user", user).Info("User added to WebSocket connection map")

	for {
		logEntry.WithField("userID", userID).Info("Starting message read loop")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logEntry.WithError(err).Error("Error reading WebSocket message")
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		logEntry.WithFields(logrus.Fields{
			"message": string(msg),
			"user":    user,
		}).Info("Received WebSocket message")

		ch.helper.SendMessageToUser(User, msg, user)
		logEntry.WithField("message", string(msg)).Info("Message sent to user")
	}
}

// GetChat handles the HTTP request to retrieve chat details.
//
// @Summary Retrieve chat details
// @Description Retrieves chat details based on the provided request
// @Tags Chat
// @Accept json
// @Produce json
// @Param body body models.ChatRequest true "Chat request details"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully retrieved chat details"
// @Failure 400 {object} response.Response{} "Details not in correct format" or "User ID not found in JWT claims" or "Failed to get chat details"
// @Router /employer/chats [post]
func (ch *ChatHandler) GetChat(c *gin.Context) {
	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID := strconv.Itoa(int(userIDInterface.(int32)))

	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}

// GroupMessage handles WebSocket group chat messages.
//
// @Summary Process WebSocket group chat messages
// @Description Processes WebSocket messages for group chat based on the provided group ID
// @Tags Chat
// @Accept json
// @Produce json
// @Param groupID path string true "Group ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "WebSocket connection established"
// @Failure 400 {object} response.Response{} "Missing Authorization header" or "Invalid token" or "Websocket Connection Issue" or "Error reading WebSocket message" or "Details not in correct format"
// @Router /group/:groupID/chat [get]
func (ch *ChatHandler) GroupMessage(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GroupMessage")
	logEntry.Info("Processing GroupMessage request")

	tokenString := c.Request.Header.Get("Authorization")
	logEntry.Info("Extracted Authorization header")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		logEntry.Error("Missing Authorization header")
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])
	logEntry.WithFields(logrus.Fields{
		"userID": userID,
		"error":  err,
	}).Info("Validated token")

	if err != nil {
		logEntry.WithError(err).Error("Invalid token")
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	groupID := c.Param("groupID")

	logEntry.Info("Upgrading to WebSocket connection")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logEntry.WithError(err).Error("WebSocket Connection Issue")
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		groupKey := groupID + "_" + strconv.Itoa(int(userID.Id))
		delete(User, groupKey)
		conn.Close()
		logEntry.Info("WebSocket connection closed and user removed from group chat")
	}()

	user := strconv.Itoa(int(userID.Id))
	groupKey := groupID + "_" + user
	User[groupKey] = conn
	logEntry.WithFields(logrus.Fields{
		"user":    user,
		"groupID": groupID,
	}).Info("User added to group chat")

	for {
		logEntry.Info("Starting message read loop")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logEntry.WithError(err).Error("Error reading WebSocket message")
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		logEntry.WithFields(logrus.Fields{
			"message": string(msg),
			"user":    user,
			"groupID": groupID,
		}).Info("Received WebSocket message")

		ch.helper.SendMessageToGroup(User, msg, groupID, user)
		logEntry.WithField("message", string(msg)).Info("Message sent to group")
	}
}
