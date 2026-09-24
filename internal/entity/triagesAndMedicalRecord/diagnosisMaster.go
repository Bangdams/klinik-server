package entity

import (
	"time"

	"github.com/google/uuid"
)

type DiagnosisMaster struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Code        string    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
