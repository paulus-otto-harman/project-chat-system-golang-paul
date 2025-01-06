package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRoomRepository(db *gorm.DB, log *zap.Logger) *RoomRepository {
	return &RoomRepository{db: db, log: log}
}
