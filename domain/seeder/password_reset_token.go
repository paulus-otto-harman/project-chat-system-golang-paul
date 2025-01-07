package seeder

import (
	"github.com/google/uuid"
	"homework/domain"
	"homework/helper"
	"time"
)

func PasswordResetTokenSeed() []domain.PasswordResetToken {
	testUuid, _ := uuid.Parse("1052c225-9a44-4f61-a340-040ef44e8022")
	return []domain.PasswordResetToken{
		{
			UserID:    2,
			Email:     "user2@mail.com",
			Otp:       "111111",
			CreatedAt: time.Now(),
			ExpiredAt: time.Now().Add(time.Hour),
		},
		{
			ID:        testUuid,
			UserID:    2,
			Email:     "user4@mailinator.com",
			Otp:       "222222",
			CreatedAt: time.Now(),
			ExpiredAt: helper.DateTime("2025-12-31 23:59:59"),
		},
		{
			UserID:      2,
			Email:       "user2@mail.com",
			Otp:         "333333",
			CreatedAt:   time.Now(),
			ExpiredAt:   time.Now().Add(time.Hour),
			ValidatedAt: helper.Ptr(time.Now()),
		},
		{
			UserID:          2,
			Email:           "user2@mail.com",
			Otp:             "444444",
			CreatedAt:       time.Now(),
			ExpiredAt:       time.Now().Add(time.Hour),
			ValidatedAt:     helper.Ptr(time.Now()),
			PasswordResetAt: helper.Ptr(time.Now()),
		},
	}
}
