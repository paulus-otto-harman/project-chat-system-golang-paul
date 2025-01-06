package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"homework/database"
)

type ChatRepository struct {
	db     *gorm.DB
	cacher database.Cacher
	log    *zap.Logger
}

func NewChatRepository(db *gorm.DB, cacher database.Cacher, log *zap.Logger) *ChatRepository {
	return &ChatRepository{db: db, cacher: cacher, log: log}
}
