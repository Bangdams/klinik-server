package entity

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	ClinicId     uuid.UUID `gorm:"not null;unique"`
	DepartmentId uuid.UUID `gorm:"not null;unique"`
	Code         string    `gorm:"not null;unique"`
	Name         string    `gorm:"not null"`
	floor        string    `gorm:"not null"`
	status       string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
