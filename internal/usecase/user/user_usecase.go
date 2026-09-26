package user

import (
	"context"
	"klinikserver/internal/model"
	"klinikserver/internal/model/converter"
	"klinikserver/internal/repository"
	"klinikserver/internal/util"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	FindByIdForUpdate(ctx context.Context, userId string) (*model.UserResponse, error)
	Create(ctx context.Context, request *model.UserRequest) (*model.UserResponse, error)
	Delete(ctx context.Context, userId string) error
	FindAll(ctx context.Context, userId string, order string, page int, limit int, sortBy string) (*[]model.UserResponse, *int, *int, *int, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, *string, error)
	Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error)
}

type UserUsecaseImpl struct {
	UserRepo repository.UserRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserUsecase(userRepo repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		UserRepo: userRepo,
		DB:       DB,
		Validate: validate,
	}
}

// Login implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, *string, error) {
	user := &model.LoginResult{
		Username: request.Username,
	}

	if err := userUsecase.UserRepo.Login(userUsecase.DB.WithContext(ctx), user); err != nil {
		log.Println("Login failed, username not found:", request.Username)
		return nil, nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		log.Println("Login failed, invalid password for:", request.Username)
		return nil, nil, fiber.ErrUnauthorized
	}

	token, err := util.GenerateTokenLogin(user)
	if err != nil {
		log.Printf("Failed to generate token for user %s: %v\n", request.Username, err)
		return nil, nil, fiber.ErrInternalServerError
	}

	log.Println("Success login:", request.Username)

	return converter.LoginUserToResponse(user), &token, nil
}
