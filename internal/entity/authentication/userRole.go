package entity

import "github.com/google/uuid"

type UserRole struct {
	UserId uuid.UUID `gorm:"not null;uniqueIndex:idx_user_role"`
	RoleId uuid.UUID `gorm:"not null;uniqueIndex:idx_user_role"`

	User User `gorm:"foreignKey:UserId;references:ID"`
	Role Role `gorm:"foreignKey:RoleId;references:ID"`
}
