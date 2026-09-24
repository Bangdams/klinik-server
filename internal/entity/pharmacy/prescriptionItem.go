package entity

import (
	"time"

	"github.com/google/uuid"
)

type PrescriptionItem struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	PrescriptionId uuid.UUID `gorm:"not null"`
	MedicineId     uuid.UUID `gorm:"not null"`
	Dosage         string    `gorm:"not null"`
	Frequency      string    `gorm:"not null"`
	Duration       string    `gorm:"not null"`
	Unit           string    `gorm:"not null"`
	Route          string    `gorm:"not null"`
	Instruction    string    `gorm:"not null"`
	Notes          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
