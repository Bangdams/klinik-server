package entity

import (
	"time"

	"github.com/google/uuid"
)

type doctorSchedule struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	DoctorId     uuid.UUID `gorm:"not null;unique"`
	DepartmentId uuid.UUID `gorm:"not null;unique"`
	dayOfWeek    uint      `gorm:"not null;unique"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
	SlotDuration uint      `gorm:"not null"`
	MaxPatients  uint      `gorm:"not null"`
	Status       string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
