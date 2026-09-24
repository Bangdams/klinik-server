package entity

import (
	"time"

	"github.com/google/uuid"
)

type Clinic struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Code      string    `gorm:"not null;unique"`
	Name      string    `gorm:"not null"`
	Address   string
	Phone     string
	Email     string `gorm:"unique"`
	Timezone  string
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Departments []Department `gorm:"foreignKey:ClinicId"`
}
