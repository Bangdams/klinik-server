package entity

import (
	"time"

	"github.com/google/uuid"
)

type DispensationItem struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	DispensationId     uuid.UUID `gorm:"not null"`
	PrescriptionItemId uuid.UUID `gorm:"not null"`
	BatchId            uuid.UUID `gorm:"not null"`
	Quantity           uint      `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
