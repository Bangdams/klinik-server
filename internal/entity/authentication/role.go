package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"not null;unique"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserRole []UserRole `gorm:"foreignKey:RoleId;references:ID"`
}
