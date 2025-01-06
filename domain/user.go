package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" example:"user1@mail.com" json:"email"`
	Password  string         `example:"password" json:"password"`
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	PasswordResetTokens []PasswordResetToken `gorm:"foreignKey:Email;references:Email" json:"-"`
}
