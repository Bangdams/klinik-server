package util

import (
	"fmt"
	"klinikserver/internal/model"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenLogin(request *model.LoginResult) (string, error) {
	var token model.TokenPyload
	duration := os.Getenv("DURATION_JWT_TOKEN")
	lifeTime, _ := strconv.Atoi(duration)

	now := time.Now()
	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "KlinikQR",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour * time.Duration(lifeTime))),
	}

	token.UserID = request.ID
	token.Username = request.Username
	token.FullName = request.FullName

	for _, v := range request.RoleName {
		token.Role = append(token.Role, v)
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return _token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GenerateTokenAttendance(startedAt, expiresAt time.Time, sessionId uint) (string, error) {
	var token model.QrTokenPyload

	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "AbsenQr",
		IssuedAt:  jwt.NewNumericDate(startedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token.SessionId = sessionId

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	Token_Result, err := _token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		fmt.Println("Error Generate Token Attendance:", err)
		return "", fiber.ErrBadRequest
	}

	return Token_Result, err
}
