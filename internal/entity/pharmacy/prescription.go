package entity

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	VisitId            uuid.UUID `gorm:"not null"`
	DoctorId           uuid.UUID `gorm:"not null"`
	PrescriptionNumber string    `gorm:"not null"`
	Status             string    `gorm:"not null"`
	Notes              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
