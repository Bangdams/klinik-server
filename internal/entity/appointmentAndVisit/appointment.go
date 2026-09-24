package entity

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID                uuid.UUID `gorm:"primaryKey"`
	ClinicId          uuid.UUID `gorm:"not null"`
	PatientId         uuid.UUID `gorm:"not null"`
	DoctorId          uuid.UUID `gorm:"not null"`
	DepartmentId      uuid.UUID `gorm:"not null"`
	ScheduleId        uuid.UUID `gorm:"not null"`
	AppointmentDate   time.Time `gorm:"not null"`
	AppointmentNumber string    `gorm:"not null"`
	Status            string    `gorm:"not null"`
	BookingSource     string    `gorm:"not null"`
	Notes             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CancelledAt       time.Time
}
