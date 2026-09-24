package entity

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	VisitId          uuid.UUID `gorm:"not null"`
	NurseId          uuid.UUID `gorm:"not null"`
	ChiefComplaint   string    `gorm:"not null"`
	Weight           float32   `gorm:"not null"`
	Height           float32   `gorm:"not null"`
	Temperature      float32   `gorm:"not null"`
	Systolic         uint      `gorm:"not null"`
	Diastolic        uint      `gorm:"not null"`
	HeartRate        uint      `gorm:"not null"`
	RespiratoryRate  uint      `gorm:"not null"`
	OxygenSaturation float32   `gorm:"not null"`
	PainScale        uint      `gorm:"not null"`
	Notes            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
