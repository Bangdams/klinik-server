package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	ClinicId     uuid.UUID `gorm:"not null;unique"`
	Username     string    `gorm:"not null;unique"`
	Email        string    `gorm:"unique"`
	PasswordHash string    `gorm:"not null"`
	FullName     string    `gorm:"not null"`
	Phone        string    `gorm:"unique"`
	Status       string    `gorm:"not null"`
	LastLoginAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	UserRole []UserRole `gorm:"foreignKey:UserId;references:ID"`
}
