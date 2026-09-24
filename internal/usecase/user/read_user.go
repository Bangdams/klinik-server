package user

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
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FindById implements UserUsecase.
func (userUsecase *UserUsecaseImpl) FindByIdForUpdate(ctx context.Context, userId string, role string) (*model.UserResponse, error) {
	user := &entity.User{
		ID: uuid.MustParse(userId),
		// Role: role,
	}

	err := userUsecase.UserRepo.FindByIdForUpdate(userUsecase.DB.WithContext(ctx), user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "User data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error findbyid user : ", err)

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error findbyid user : ", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

// FindAll implements UserUsecase.
func (userUsecase *UserUsecaseImpl) FindAll(ctx context.Context, userId string, role string, order string, page int, limit int, sortBy string) (*[]model.UserResponse, *int, *int, *int, error) {
	var users = &[]entity.User{}

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

	totalRecords, err := userUsecase.UserRepo.FindAllForPagging(userUsecase.DB.WithContext(ctx), uuid.MustParse(userId), role, limit, offset, order, sortBy, users)
	if err != nil {
		log.Println("failed when find all repo user : ", err)
		return nil, nil, nil, nil, fiber.ErrInternalServerError
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if totalPages == 0 {
		totalPages = 1
	}

	taotalRecodeInt := int(totalRecords)

	log.Println("success find all from usecase user")

	return converter.UserToResponses(users), &page, &taotalRecodeInt, &totalPages, nil
}
