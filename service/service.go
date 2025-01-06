package service

import (
	"go.uber.org/zap"
	"homework/config"
	"homework/repository"
)

type Service struct {
	Auth          AuthService
	Chat          ChatService
	Email         EmailService
	Otp           OtpService
	PasswordReset PasswordResetService
	Room          RoomService
	User          UserService
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.User, log),
		Chat:          NewChatService(repo.Chat, log),
		Email:         NewEmailService(appConfig.Email, log),
		Otp:           NewOtpService(log),
		PasswordReset: NewPasswordResetService(repo.PasswordReset, log),
		Room:          NewRoomService(repo.Room, log),
		User:          NewUserService(repo, log),
	}
}
