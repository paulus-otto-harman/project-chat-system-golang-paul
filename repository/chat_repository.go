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

func (r ChatRepository) Save(message domain.Message) error {
	return r.db.Create(&message).Error
}

func (r ChatRepository) Delete(messageId string) error {
	return r.db.Delete(&domain.Message{}, messageId).Error
}

func (r ChatRepository) All(roomId string) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("room_id = ?", roomId).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
