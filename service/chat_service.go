package service

import (
	"go.uber.org/zap"
	"homework/domain"
	"homework/repository"
	"log"
)

type ChatService interface {
	SaveMessage(message domain.Message) error
}

type chatService struct {
	repo repository.ChatRepository
	log  *zap.Logger
}

func NewChatService(repo repository.ChatRepository, log *zap.Logger) ChatService {
	return &chatService{repo, log}
}

func (s *chatService) SaveMessage(message domain.Message) error {
	log.Println("Saving message", zap.Any("message", message))
	s.repo.Save(message)
	return nil
}
