package entity

import (
	"time"

	"github.com/google/uuid"
)

type VisitProcedure struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	VisitId     uuid.UUID `gorm:"not null"`
	ProcedureId uuid.UUID `gorm:"not null"`
	DoctorId    uuid.UUID `gorm:"not null"`
	Quantity    uint      `gorm:"not null"`
	Price       uint      `gorm:"not null"`
	Notes       string
	CreatedAt   time.Time
}
