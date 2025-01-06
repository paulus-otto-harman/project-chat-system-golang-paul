package repository

import (
	"homework/config"
	"homework/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher, log),
		PasswordReset: *NewPasswordResetRepository(db, log),
		User:          *NewUserRepository(db, log),
	}
}
