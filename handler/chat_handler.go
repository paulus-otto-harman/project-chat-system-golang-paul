package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"homework/database"
	"homework/service"
)

type ChatController struct {
	cacher  database.Cacher
	service service.ChatService
	logger  *zap.Logger
}

func NewChatController(service service.ChatService, logger *zap.Logger, cacher database.Cacher) *ChatController {
	return &ChatController{cacher, service, logger}
}

func (ctrl *ChatController) Websocket(c *gin.Context) {

}

func (ctrl *ChatController) All(c *gin.Context) {

}

func (ctrl *ChatController) Delete(c *gin.Context) {

}
