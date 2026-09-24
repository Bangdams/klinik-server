package entity

import (
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	ClinicId        uuid.UUID `gorm:"not null;unique"`
	UserId          uuid.UUID `gorm:"not null;unique"`
	MedicalRecordNo string    `gorm:"not null;unique"`
	Nik             string    `gorm:"not null"`
	FullName        string    `gorm:"not null"`
	BirthDate       time.Time `gorm:"not null"`
	BirthPlace      string    `gorm:"not null"`
	Gender          string    `gorm:"not null"`
	BloodType       string    `gorm:"not null"`
	Phone           string    `gorm:"not null;unique"`
	Email           string    `gorm:"not null;unique"`
	Occupation      string    `gorm:"not null"`
	MaritalStatus   string    `gorm:"not null"`
	Status          string    `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
