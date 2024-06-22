package handler

import (
	logging "HireoGateWay/Logging"
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"
	"os"
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
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/Hireo_gateway.log")
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// WebSocket
func (ch *ChatHandler) EmployerMessage(c *gin.Context) {
	fmt.Println("++== call hit in message funtion")
	tokenString := c.Request.Header.Get("Authorization")
	ch.Logger.Info("Extracted Authorization header")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		ch.Logger.Error("Missing Authorization header")
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])
	ch.Logger.WithFields(logrus.Fields{
		"userID": userID,
		"error":  err,
	}).Info("Validated token")

	if err != nil {
		ch.Logger.WithError(err).Error("Invalid token")
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	ch.Logger.Info("Upgrading to WebSocket connection")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ch.Logger.WithError(err).Error("WebSocket Connection Issue")
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		delete(User, strconv.Itoa(int(userID.Id)))
		conn.Close()
		ch.Logger.Info("WebSocket connection closed and user removed from map")
	}()

	user := strconv.Itoa(int(userID.Id))
	User[user] = conn
	ch.Logger.WithField("user", user).Info("User added to WebSocket connection map")

	for {
		ch.Logger.WithField("userID", userID).Info("Starting message read loop")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			ch.Logger.WithError(err).Error("Error reading WebSocket message")
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		ch.Logger.WithFields(logrus.Fields{
			"message": string(msg),
			"user":    user,
		}).Info("Received WebSocket message")

		ch.helper.SendMessageToUser(User, msg, user)
		ch.Logger.WithField("message", string(msg)).Info("Message sent to user")
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

	ch.Logger.Info("Processing GroupMessage request")
	tokenString := c.Request.Header.Get("Authorization")
	ch.Logger.Info("Extracted Authorization header")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		ch.Logger.Error("Missing Authorization header")
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])
	ch.Logger.WithFields(logrus.Fields{
		"userID": userID,
		"error":  err,
	}).Info("Validated token")

	if err != nil {
		ch.Logger.WithError(err).Error("Invalid token")
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	groupID := c.Param("groupID")

	ch.Logger.Info("Upgrading to WebSocket connection")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ch.Logger.WithError(err).Error("WebSocket Connection Issue")
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		groupKey := groupID + "_" + strconv.Itoa(int(userID.Id))
		delete(User, groupKey)
		conn.Close()
		ch.Logger.Info("WebSocket connection closed and user removed from group chat")
	}()

	user := strconv.Itoa(int(userID.Id))
	groupKey := groupID + "_" + user
	User[groupKey] = conn
	ch.Logger.WithFields(logrus.Fields{
		"user":    user,
		"groupID": groupID,
	}).Info("User added to group chat")

	for {
		ch.Logger.Info("Starting message read loop")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			ch.Logger.WithError(err).Error("Error reading WebSocket message")
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		ch.Logger.WithFields(logrus.Fields{
			"message": string(msg),
			"user":    user,
			"groupID": groupID,
		}).Info("Received WebSocket message")

		ch.helper.SendMessageToGroup(User, msg, groupID, user)
		ch.Logger.WithField("message", string(msg)).Info("Message sent to group")
	}
}
