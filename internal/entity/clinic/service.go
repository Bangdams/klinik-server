package entity

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClinicId    uuid.UUID `gorm:"not null;unique"`
	Code        string    `gorm:"not null;unique"`
	Name        string    `gorm:"not null"`
	Description string
	Price       int `gorm:"not null"`
	Status      int `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
