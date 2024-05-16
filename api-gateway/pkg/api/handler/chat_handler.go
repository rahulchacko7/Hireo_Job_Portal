package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
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

func (ch *JobHandler) EmployerMessage(c *gin.Context) {
	fmt.Println("message called")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userID, ok := c.Get("user_id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID in JWT claims is not a string", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer delete(User, strconv.Itoa(userID.(int)))
	defer conn.Close()
	user := strconv.Itoa(userID.(int))
	User[user] = conn

	for {
		fmt.Println("loop starts", userID, User)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		ch.helper.SendMessageToUser(User, msg, user)
	}
}
