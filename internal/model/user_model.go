package model

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	FullName  string    `json:"fullname" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt string    `json:"created_at" validate:"required"`
}

type UserRequest struct {
	ClinicId     string   `json:"clinic_id" validate:"required"`
	Username     string   `json:"username" validate:"required"`
	Email        string   `json:"email"`
	PasswordHash string   `json:"password" validate:"required"`
	FullName     string   `json:"fullname" validate:"required"`
	Phone        string   `json:"phone" validate:"required"`
	Status       string   `json:"status" validate:"required"`
	RoleId       []string `json:"role_id" validate:"required"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password"`
	Role     string    `json:"role" validate:"required"`

	// CoachRequest   *CoachRequest   `json:"coach" validate:"required_if=Role coach"`
	// StudentRequest *StudentRequest `json:"student" validate:"required_if=Role student"`
}
