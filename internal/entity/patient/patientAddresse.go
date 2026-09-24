package entity

import (
	"time"

	"github.com/google/uuid"
)

type PatientAddresses struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	PatientId  uuid.UUID `gorm:"not null;unique"`
	Label      string    `gorm:"not null"`
	Address    string    `gorm:"not null"`
	Province   string    `gorm:"not null"`
	City       string    `gorm:"not null"`
	District   string    `gorm:"not null"`
	Village    string    `gorm:"not null"`
	PostalCode string    `gorm:"not null"`
	IsPrimary  bool      `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
