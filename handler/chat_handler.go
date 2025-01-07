package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"homework/database"
	"homework/domain"
	"homework/helper"
	"homework/service"
	"log"
	"net/http"
)

type ChatController struct {
	service service.ChatService
	cacher  database.Cacher
	logger  *zap.Logger
}

func NewChatController(service service.ChatService, logger *zap.Logger, cacher database.Cacher) *ChatController {
	return &ChatController{service, cacher, logger}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins. Modify as per your security needs.
		return true
	},
}

func (ctrl *ChatController) Websocket(c *gin.Context) {
	if c.GetString("user-id") == "" {
		BadResponse(c, "unauthorized", http.StatusUnauthorized)
		return
	}
	userId, _ := helper.Uint(c.GetString("user-id"))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	roomId, _ := helper.Uint(c.Param("id"))
	room := fmt.Sprintf("room:%d", roomId)
	subscriber := ctrl.cacher.Subscribe(room)
	message := domain.Message{RoomID: roomId, SenderID: userId}

	go func() {
		for {
			payload, _ := subscriber.ReceiveMessage(context.Background())
			conn.WriteMessage(websocket.TextMessage, []byte(payload.Payload))
		}
	}()

	for {
		_, msg, _ := conn.ReadMessage()
		message.Content = string(msg)

		ctrl.service.SaveMessage(message)
		ctrl.cacher.Publish(room, message.Content)
	}
}

func (ctrl *ChatController) All(c *gin.Context) {
	id := c.Param("id")
	messages, err := ctrl.service.GetMessages(id)
	if err != nil {
		BadResponse(c, "failed to get all messages", http.StatusInternalServerError)
	}
	GoodResponseWithData(c, "message successfully retrieved", http.StatusOK, messages)
}

func (ctrl *ChatController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.DeleteMessage(id); err != nil {
		BadResponse(c, "failed to delete message id "+id, http.StatusInternalServerError)
	}
	GoodResponseWithData(c, "message successfully deleted", http.StatusOK, nil)
}
