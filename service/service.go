package service

import (
	"go.uber.org/zap"
	"homework/config"
	"homework/repository"
)

type Service struct {
	Auth          AuthService
	Email         EmailService
	Otp           OtpService
	PasswordReset PasswordResetService
	User          UserService
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.User, log),
		Email:         NewEmailService(appConfig.Email, log),
		Otp:           NewOtpService(log),
		PasswordReset: NewPasswordResetService(repo.PasswordReset, log),
		User:          NewUserService(repo, log),
	}
}
