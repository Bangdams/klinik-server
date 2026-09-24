package entity

import (
	"time"

	"github.com/google/uuid"
)

type PatientEmergencyContacts struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	PatientId    uuid.UUID `gorm:"not null;unique"`
	Name         string    `gorm:"not null"`
	Relationship string    `gorm:"not null"`
	Phone        string    `gorm:"not null"`
	Address      string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
