package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	UserId           uuid.UUID `gorm:"not null"`
	RefreshTokenHash string    `gorm:"not null"`
	IpAddress        string    `gorm:"not null"`
	UserAgent        string    `gorm:"not null"`
	ExpiresAt        time.Time
	CreatedAt        time.Time
	RevokedAt        time.Time
}
