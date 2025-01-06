package service

import (
	"go.uber.org/zap"
	"homework/repository"
)

type ChatService interface {
}

type chatService struct {
	repo repository.ChatRepository
	log  *zap.Logger
}

func NewChatService(repo repository.ChatRepository, log *zap.Logger) ChatService {
	return &chatService{repo, log}
}
