package role

import (
	"context"
	"encoding/json"
	"errors"
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"klinikserver/internal/model/converter"
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// FindAll implements [RoleUsecase].
func (roleUsecase *RoleUsecaseImpl) FindAll(ctx context.Context, name string, order string, page int, limit int) (*[]model.RoleResponse, *int, *int, *int, error) {
	var roles = &[]entity.Role{}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	if order == "" {
		order = "DESC"
	}

	totalRecords, err := roleUsecase.RoleRepo.FindAll(roleUsecase.DB.WithContext(ctx), name, limit, offset, order, roles)
	if err != nil {
		log.Println("failed when find all repo role : ", err)
		return nil, nil, nil, nil, fiber.ErrInternalServerError
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if totalPages == 0 {
		totalPages = 1
	}

	taotalRecodeInt := int(totalRecords)

	log.Println("success find all from usecase role")

	return converter.RoleToResponses(roles), &page, &taotalRecodeInt, &totalPages, nil
}

// FindByName implements [RoleUsecase].
func (roleUsecase *RoleUsecaseImpl) FindByName(ctx context.Context, name string) (*model.RoleResponse, error) {
	role := &entity.Role{
		Name: name,
	}

	err := roleUsecase.RoleRepo.FindByName(roleUsecase.DB.WithContext(ctx), role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "Role data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error find role by name : ", err)

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error find role by name : ", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleToResponse(role), nil
}
