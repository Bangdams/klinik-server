package entity

import (
	"time"

	"github.com/google/uuid"
)

type PatientQr struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	PatientId uuid.UUID `gorm:"not null;unique"`
	tokenHash string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	IssuedAt  time.Time
	ExpiresAt time.Time
	RevokedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
