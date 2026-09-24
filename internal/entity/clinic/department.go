package entity

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClinicId    uuid.UUID `gorm:"not null;unique"`
	Code        string    `gorm:"unique"`
	Name        string    `gorm:"not null"`
	Description string
	Status      string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Clinic Clinic `gorm:"foreignKey:ClinicId;references:ID"`
}
