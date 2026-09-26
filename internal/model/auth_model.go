package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	RoleName []string  `json:"role"`
}

type LoginResult struct {
	ID           uuid.UUID `json:"user_id"`
	Username     string    `json:"username"`
	FullName     string    `json:"full_name"`
	PasswordHash string
	RoleName     []string `json:"role"`
}

type LoginRow struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	FullName     string
	RoleName     string
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
	Role     []string  `json:"role"`
	TokenJwt
}
