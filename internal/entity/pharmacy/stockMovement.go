package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockMovement struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	MedicineId    uuid.UUID `gorm:"not null"`
	BatchId       uuid.UUID `gorm:"not null"`
	Type          string    `gorm:"not null"`
	Quantity      uint      `gorm:"not null"`
	ReferenceType string    `gorm:"not null"`
	ReferenceId   uuid.UUID `gorm:"not null"`
	Notes         string
	CreatedBy     uuid.UUID `gorm:"not null"`
	CreatedAt     time.Time
}
