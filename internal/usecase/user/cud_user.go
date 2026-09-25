package user

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
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Create implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Create(ctx context.Context, request *model.UserRequest) (*model.UserResponse, error) {
	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := userUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create user : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to generate password")
		return nil, fiber.ErrInternalServerError
	}

	roles := []entity.UserRole{}

	for _, v := range request.RoleId {
		roles = append(roles, entity.UserRole{
			RoleId: uuid.MustParse(v),
		})
	}

	user := &entity.User{
		ClinicId:     uuid.MustParse(request.ClinicId),
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: string(password),
		FullName:     request.FullName,
		Phone:        request.Phone,
		Status:       request.Status,
		UserRole:     roles,
	}

	if err := userUsecase.UserRepo.FindByUsername(tx, user); err == nil {
		errorResponse.Message = "Duplicate entry"
		errorResponse.Details = []string{"Username already exists in the database."}

		jsonString, _ := json.Marshal(errorResponse)

		return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
	}

	err = userUsecase.UserRepo.Create(tx, user)
	if err != nil {
		log.Println("failed when create repo user : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase user")

	return converter.UserToResponse(user), nil
}

// Delete implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Delete(ctx context.Context, userId string) error {
	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := &entity.User{}
	user.ID = uuid.MustParse(userId)

	err := userUsecase.UserRepo.FindById(tx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "User data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete user : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete user : ", err)
		return fiber.ErrInternalServerError
	}

	err = userUsecase.UserRepo.Delete(tx, user)
	if err != nil {
		log.Println("failed when delete repo user : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase user")

	return nil
}

// Update implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := userUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error update user : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	user := &entity.User{
		ID: request.ID,
	}

	err = userUsecase.UserRepo.FindByIdForUpdate(tx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse.Message = "User data was not found"
			errorResponse.Details = []string{}

			jsonString, _ := json.Marshal(errorResponse)
			log.Println("Data not found")

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error find by user id : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("failed to generate password")
			return nil, fiber.ErrInternalServerError
		}

		user.PasswordHash = string(password)
	}

	user.Username = request.Username

	// switch user.Role {
	// case "coach":
	// 	user.Coach = entity.Coach{
	// 		Nip:      request.CoachRequest.Nip,
	// 		FullName: request.CoachRequest.FullName,
	// 	}
	// case "student":
	// 	user.Student = entity.Student{
	// 		Nis:         request.StudentRequest.Nis,
	// 		FullName:    request.StudentRequest.FullName,
	// 		Address:     request.StudentRequest.Address,
	// 		PhoneNumber: request.StudentRequest.PhoneNumber,
	// 	}
	// }

	err = userUsecase.UserRepo.Update(tx, user)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		log.Println("failed when update repo user : ", err)

		var errorField string
		parts := strings.Split(mysqlErr.Message, "'")
		if len(parts) > 2 {
			errorField = parts[1]
		}

		if mysqlErr.Number == 1062 {
			errorResponse.Message = "Duplicate entry"
			errorResponse.Details = []string{errorField + " already exists in the database."}

			jsonString, _ := json.Marshal(errorResponse)

			return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
		}

		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase user")

	return converter.UserToResponse(user), nil
}
