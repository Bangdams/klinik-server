package entity

import (
	"time"

	"github.com/google/uuid"
)

type Triage struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	VisitId    uuid.UUID `gorm:"not null"`
	DoctorId   uuid.UUID `gorm:"not null"`
	Subjective string    `gorm:"not null"`
	Objective  string    `gorm:"not null"`
	Assessment string    `gorm:"not null"`
	Plan       string    `gorm:"not null"`
	Notes      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
