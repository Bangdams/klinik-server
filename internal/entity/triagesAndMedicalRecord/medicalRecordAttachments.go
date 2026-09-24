package entity

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecordAttachment struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	MedicalRecordId uuid.UUID `gorm:"not null"`
	FileName        string    `gorm:"not null"`
	FilePath        string    `gorm:"not null"`
	MimeType        string    `gorm:"not null"`
	FileSize        uint      `gorm:"not null"`
	UploadedBy      uuid.UUID `gorm:"not null"`
	CreatedAt       time.Time
}
