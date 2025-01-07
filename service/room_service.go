package service

import (
	"go.uber.org/zap"
	"homework/domain"
	"homework/repository"
)

type RoomService interface {
	SaveRoom(room *domain.Room) error
}

type roomService struct {
	repo repository.RoomRepository
	log  *zap.Logger
}

func NewRoomService(repo repository.RoomRepository, log *zap.Logger) RoomService {
	return &roomService{repo, log}
}

func (s *roomService) SaveRoom(room *domain.Room) error {
	return s.repo.Save(room)
}
