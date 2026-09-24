package entity

import (
	"time"

	"github.com/google/uuid"
)

type VisitQr struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	VisitId   uuid.UUID `gorm:"not null;unique"`
	tokenHash string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	ExpiresAt time.Time
	ScannedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
