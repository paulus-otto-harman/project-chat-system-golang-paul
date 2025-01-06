package handler

import (
	"go.uber.org/zap"
	"homework/service"
)

type ChatController struct {
	service service.UserService
	logger  *zap.Logger
}

func NewChatController(service service.UserService, logger *zap.Logger) *ChatController {
	return &ChatController{service: service, logger: logger}
}
