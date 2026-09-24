package entity

import (
	"time"

	"github.com/google/uuid"
)

type MedicineBatche struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	MedicineId    uuid.UUID `gorm:"not null"`
	BatchNumber   string    `gorm:"not null"`
	ExpiredAt     time.Time `gorm:"not null"`
	PurchasePrice uint      `gorm:"not null"`
	SellingPrice  uint      `gorm:"not null"`
	Quantity      uint      `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
