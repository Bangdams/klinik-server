package util

import (
	"klinikserver/internal/model"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(request *model.LoginResult) (string, error) {
	var token model.TokenPyload
	duration := os.Getenv("DURATION_JWT_ACCESS_TOKEN")
	lifeTime, _ := strconv.Atoi(duration)

	now := time.Now()
	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "KlinikQR",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * time.Duration(lifeTime))),
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

func GenerateRefreshToken(request *model.LoginResult) (string, error) {
	var token model.TokenPyload
	duration := os.Getenv("DURATION_JWT_REFRESH_TOKEN")
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
