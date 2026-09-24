package entity

import (
	"time"

	"github.com/google/uuid"
)

type Medicine struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Status      string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
