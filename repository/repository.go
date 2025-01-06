package repository

import (
	"homework/config"
	"homework/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	Chat          ChatRepository
	PasswordReset PasswordResetRepository
	Room          RoomRepository
	User          UserRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher, log),
		Chat:          *NewChatRepository(db, cacher, log),
		PasswordReset: *NewPasswordResetRepository(db, log),
		Room:          *NewRoomRepository(db, log),
		User:          *NewUserRepository(db, log),
	}
}
