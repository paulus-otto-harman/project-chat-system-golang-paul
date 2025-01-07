package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"homework/domain"
)

type RoomRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRoomRepository(db *gorm.DB, log *zap.Logger) *RoomRepository {
	return &RoomRepository{db: db, log: log}
}

func (r RoomRepository) Save(room domain.Room) error {
	return r.db.Create(&room).Error
}
