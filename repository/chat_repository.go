package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"homework/database"
	"homework/domain"
)

type ChatRepository struct {
	db     *gorm.DB
	cacher database.Cacher
	log    *zap.Logger
}

func NewChatRepository(db *gorm.DB, cacher database.Cacher, log *zap.Logger) *ChatRepository {
	return &ChatRepository{db: db, cacher: cacher, log: log}
}

func (r ChatRepository) Save(message domain.Message) {
	r.db.Create(&message)
}
