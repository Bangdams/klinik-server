package entity

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	UserId       uuid.UUID `gorm:"primaryKey"`
	Token        string    `gorm:"not null"`
	StatusLogout uint      `gorm:"not null "`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time
	User         User `gorm:"foreignKey:user_id;references:id"`
}
