package service

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"homework/domain"
	"homework/helper"
	"homework/repository"
	"time"
)

type PasswordResetService interface {
	Create(token *domain.PasswordResetToken) error
	Validate(id uuid.UUID, token string) error
}

type passwordResetService struct {
	repo repository.PasswordResetRepository
	log  *zap.Logger
}

func NewPasswordResetService(repo repository.PasswordResetRepository, log *zap.Logger) PasswordResetService {
	return &passwordResetService{repo, log}
}

func (s *passwordResetService) Create(token *domain.PasswordResetToken) error {
	return s.repo.Create(token)
}

func (s *passwordResetService) Validate(id uuid.UUID, token string) error {
	passwordResetToken, err := s.repo.GetValidToken(&domain.PasswordResetToken{ID: id, Otp: token})
	if err != nil {
		return err
	}

	passwordResetToken.ValidatedAt = helper.Ptr(time.Now())
	return s.repo.Update(passwordResetToken)
}
