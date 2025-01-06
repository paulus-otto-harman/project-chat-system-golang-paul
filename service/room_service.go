package service

import (
	"go.uber.org/zap"
	"homework/repository"
)

type RoomService interface {
}

type roomService struct {
	repo repository.RoomRepository
	log  *zap.Logger
}

func NewRoomService(repo repository.RoomRepository, log *zap.Logger) RoomService {
	return &roomService{repo, log}
}
