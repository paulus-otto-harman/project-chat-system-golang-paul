package service

import (
	"go.uber.org/zap"
	"homework/domain"
	"homework/repository"
	"log"
)

type ChatService interface {
	SaveMessage(message domain.Message) error
	DeleteMessage(id string) error
	GetMessages(roomId string) ([]domain.Message, error)
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
	return s.repo.Save(message)
}

func (s *chatService) DeleteMessage(id string) error {
	log.Println("Delete message", zap.String("message ID", id))
	return s.repo.Delete(id)
}

func (s *chatService) GetMessages(roomId string) ([]domain.Message, error) {
	return s.repo.All(roomId)
}
