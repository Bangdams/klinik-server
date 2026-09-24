package role

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"klinikserver/internal/model/converter"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Create implements [RoleUsecase].
func (roleUsecase *RoleUsecaseImpl) Create(ctx context.Context, request *model.RoleRequest) (*model.RoleResponse, error) {
	tx := roleUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := roleUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create role : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	role := &entity.Role{
		Name:        request.Name,
		Description: request.Description,
	}

	if err := roleUsecase.RoleRepo.FindByName(tx, role); err == nil {
		errorResponse.Message = "Duplicate entry"
		errorResponse.Details = []string{"Name already exists in the database."}

		jsonString, _ := json.Marshal(errorResponse)

		return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
	}

	err = roleUsecase.RoleRepo.Create(tx, role)
	if err != nil {
		log.Println("failed when create repo role : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase role")

	return converter.RoleToResponse(role), nil
}

// Delete implements [RoleUsecase].
func (roleUsecase *RoleUsecaseImpl) Delete(ctx context.Context, name string) error {
	tx := roleUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	role := &entity.Role{}
	role.Name = name

	err := roleUsecase.RoleRepo.FindByName(tx, role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "role data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete role : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete role : ", err)
		return fiber.ErrInternalServerError
	}

	err = roleUsecase.RoleRepo.Delete(tx, role)
	if err != nil {
		log.Println("failed when delete repo role : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase role")

	return nil
}

// Update implements [RoleUsecase].
func (roleUsecase *RoleUsecaseImpl) Update(ctx context.Context, request *model.UpdateRoleRequest) (*model.RoleResponse, error) {
	tx := roleUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := roleUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error update role : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	role := &entity.Role{
		Name: request.OldName,
	}

	err = roleUsecase.RoleRepo.FindByName(tx, role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	// Update data
	role.Name = request.Name
	role.Description = request.Description

	err = roleUsecase.RoleRepo.Update(tx, role)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		log.Println("failed when update repo role : ", err)

		var errorField string
		parts := strings.Split(mysqlErr.Message, "'")
		if len(parts) > 2 {
			errorField = parts[1]
		}

		if mysqlErr.Number == 1062 {
			errorResponse.Message = "Duplicate entry"
			errorResponse.Details = []string{errorField + " role name already exists in the database."}

			jsonString, _ := json.Marshal(errorResponse)

			return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
		}

		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase role")

	return converter.RoleToResponse(role), nil
}
