package entity

import (
	"time"

	"github.com/google/uuid"
)

type Queues struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	VisitId      uuid.UUID `gorm:"not null"`
	DepartmentId uuid.UUID `gorm:"not null"`
	QueueNumber  string    `gorm:"not null"`
	QueueDate    time.Time `gorm:"not null"`
	Status       string    `gorm:"not null"`
	CalledAt     time.Time `gorm:"not null"`
	StartedAt    time.Time `gorm:"not null"`
	CompletedAt  time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
