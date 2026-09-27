package util

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckLevel(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userToken := ctx.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)

		userRoles, err := GetRoles(claims)
		if err != nil {
			return fiber.ErrForbidden
		}

		allowed := make(map[string]struct{})
		for _, role := range roles {
			allowed[role] = struct{}{}
		}

		for _, role := range userRoles {
			if _, ok := allowed[role]; ok {
				return ctx.Next()
			}
		}

		return fiber.ErrForbidden
	}
}

func GetRoles(claims jwt.MapClaims) ([]string, error) {
	rawRoles, ok := claims["role"].([]interface{})
	if !ok {
		return nil, errors.New("invalid role claim")
	}

	roles := make([]string, 0, len(rawRoles))

	for _, r := range rawRoles {
		role, ok := r.(string)
		if !ok {
			continue
		}

		roles = append(roles, role)
	}

	return roles, nil
}
