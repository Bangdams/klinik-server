package entity

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"not null;unique"`
	ClinicId  uuid.UUID `gorm:"not null;unique"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
