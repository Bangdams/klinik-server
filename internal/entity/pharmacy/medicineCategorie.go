package entity

import (
	"time"

	"github.com/google/uuid"
)

type MedicineCategorie struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	ClinicId      uuid.UUID `gorm:"not null"`
	CategoryId    uuid.UUID `gorm:"not null"`
	Code          string    `gorm:"not null"`
	Name          string    `gorm:"not null"`
	GenericName   string    `gorm:"not null"`
	Unit          string    `gorm:"not null"`
	PurchasePrice uint      `gorm:"not null"`
	SellingPrice  uint      `gorm:"not null"`
	MinimumStock  uint      `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
