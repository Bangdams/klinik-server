package entity

import (
	"time"

	"github.com/google/uuid"
)

type Visits struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	ClinicId      uuid.UUID `gorm:"not null"`
	PatientId     uuid.UUID `gorm:"not null"`
	DoctorId      uuid.UUID `gorm:"not null"`
	DepartmentId  uuid.UUID `gorm:"not null"`
	AppointmentId uuid.UUID `gorm:"not null;unique"`
	VisitNumber   string    `gorm:"not null"`
	VisitDate     time.Time `gorm:"not null"`
	VisitType     string    `gorm:"not null"`
	PaymentType   string    `gorm:"not null"`
	Status        string    `gorm:"not null"`
	Notes         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CompletedAt   time.Time
}
