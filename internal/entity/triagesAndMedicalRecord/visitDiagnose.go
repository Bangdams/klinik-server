package entity

import (
	"time"

	"github.com/google/uuid"
)

type VisitDiagnose struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	VisitId     uuid.UUID `gorm:"not null"`
	DiagnosisId uuid.UUID `gorm:"not null"`
	Type        string    `gorm:"not null"`
	Notes       string
	CreatedAt   time.Time
}
