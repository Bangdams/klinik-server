package entity

import (
	"time"

	"github.com/google/uuid"
)

type Dispensation struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	PrescriptionId uuid.UUID `gorm:"not null"`
	PharmacistId   uuid.UUID `gorm:"not null"`
	DispensedAt    time.Time `gorm:"not null"`
	Status         string    `gorm:"not null"`
	Notes          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
