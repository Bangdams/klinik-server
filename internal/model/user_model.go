package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Role     string    `json:"role" validate:"required"`
	// CoachRequest   *CoachRequest   `json:"coach" validate:"required_if=Role coach"`
	// StudentRequest *StudentRequest `json:"student" validate:"required_if=Role student"`
	CreatedAt string `json:"created_at" validate:"required"`
}

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`

	// CoachRequest   *CoachRequest   `json:"coach" validate:"required_if=Role coach"`
	// StudentRequest *StudentRequest `json:"student" validate:"required_if=Role student"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password"`
	Role     string    `json:"role" validate:"required"`

	// CoachRequest   *CoachRequest   `json:"coach" validate:"required_if=Role coach"`
	// StudentRequest *StudentRequest `json:"student" validate:"required_if=Role student"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Role     string    `json:"role"`
}

type TokenJwt struct {
	jwt.RegisteredClaims
}

type QrTokenPyload struct {
	SessionId uint `json:"session_id"`
	TokenJwt
}

type TokenPyload struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Role     string    `json:"role"`
	TokenJwt
}
