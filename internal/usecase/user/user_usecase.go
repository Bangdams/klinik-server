package user

import (
	"context"
	"errors"
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"klinikserver/internal/model/converter"
	"klinikserver/internal/repository"
	"klinikserver/internal/util"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	FindByIdForUpdate(ctx context.Context, userId string) (*model.UserResponse, error)
	FindAll(ctx context.Context, userId string, order string, page int, limit int, sortBy string) (*[]model.UserResponse, *int, *int, *int, error)
	Login(ctx context.Context, request *model.LoginRequest, requestRefreshTokenUser string) (*model.LoginResponse, *string, error)
	Logout(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, refreshToken string) (*model.LoginResponse, error)
	Create(ctx context.Context, request *model.UserRequest) (*model.UserResponse, error)
	Delete(ctx context.Context, userId string) error
	Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error)
}

type UserUsecaseImpl struct {
	UserRepo         repository.UserRepository
	RefreshTokenRepo repository.RefreshTokenRepository
	DB               *gorm.DB
	Validate         *validator.Validate
}

func NewUserUsecase(userRepo repository.UserRepository, refreshRepo repository.RefreshTokenRepository, DB *gorm.DB, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		UserRepo:         userRepo,
		RefreshTokenRepo: refreshRepo,
		DB:               DB,
		Validate:         validate,
	}
}

// Refresh implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Refresh(ctx context.Context, refreshToken string) (*model.LoginResponse, error) {
	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	claims, err := util.ParseToken(refreshToken, []byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	userId := claims["user_id"].(string)
	roles, err := util.GetRoles(claims)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusForbidden, "Invalid Role Claim")
	}

	request := model.LoginResult{
		ID:       uuid.MustParse(userId),
		Username: claims["username"].(string),
		FullName: claims["full_name"].(string),
		RoleName: roles,
	}

	if err := userUsecase.RefreshTokenRepo.CheckStatusLogout(tx, request.ID); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	newAccessToken, _ := util.GenerateAccessToken(&request)

	log.Println("success create access token")

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.LoginUserToResponse(newAccessToken), nil
}

// Login implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Login(ctx context.Context, request *model.LoginRequest, requestRefreshTokenUser string) (*model.LoginResponse, *string, error) {
	_, err := util.ParseToken(requestRefreshTokenUser, []byte(os.Getenv("SECRET_KEY")))
	if err == nil {
		return nil, nil, fiber.NewError(fiber.StatusBadRequest, "Refresh token still valid")
	}

	loginResult := model.LoginResult{
		Username: request.Username,
	}

	if err := userUsecase.UserRepo.Login(userUsecase.DB.WithContext(ctx), &loginResult); err != nil {
		log.Println("Login failed, username not found:", request.Username)
		return nil, nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginResult.PasswordHash), []byte(request.Password)); err != nil {
		log.Println("Login failed, invalid password for:", request.Username)
		return nil, nil, fiber.ErrUnauthorized
	}

	accessToken, err := util.GenerateAccessToken(&loginResult)
	if err != nil {
		log.Println("Failed to generate token jwt")
		return nil, nil, fiber.ErrInternalServerError
	}

	refreshToken, err := util.GenerateRefreshToken(&loginResult)
	if err != nil {
		log.Println("Failed to generate token jwt")
		return nil, nil, fiber.ErrInternalServerError
	}

	duration := os.Getenv("DURATION_JWT_REFRESH_TOKEN")
	lifeTime, _ := strconv.Atoi(duration)
	now := time.Now()

	requestRefreshToken := &entity.RefreshToken{
		UserId:       loginResult.ID,
		StatusLogout: 0,
		Token:        refreshToken,
		ExpiresAt:    now.Add(time.Minute * time.Duration(lifeTime)),
	}

	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := userUsecase.RefreshTokenRepo.FindById(tx, loginResult.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			log.Println("Creating new refresh token in database")

			if err := userUsecase.RefreshTokenRepo.Create(tx, requestRefreshToken); err != nil {
				return nil, nil, err
			}

		} else {
			return nil, nil, err
		}
	} else {

		log.Println("Updating existing refresh token in database")

		if err := userUsecase.RefreshTokenRepo.Update(tx, requestRefreshToken); err != nil {
			return nil, nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, nil, fiber.ErrInternalServerError
	}

	log.Println("Success login:", request.Username)

	return converter.LoginUserToResponse(accessToken), &refreshToken, nil
}

// Logout implements UserUsecase.
func (userUsecase *UserUsecaseImpl) Logout(ctx context.Context, refreshToken string) error {
	tx := userUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	claims, err := util.ParseToken(refreshToken, []byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return fiber.ErrUnauthorized
	}

	userId := claims["user_id"].(string)

	if err := userUsecase.RefreshTokenRepo.FindById(tx, uuid.MustParse(userId)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Logout failed because the user has not logged in.")
			return fiber.NewError(fiber.StatusBadRequest, "User has not logged in")
		} else {
			log.Println("Error RefreshToken findbyid:", err)
			return fiber.ErrInternalServerError
		}
	} else {
		log.Println("Logout successful.")
		userUsecase.RefreshTokenRepo.Update(tx, &entity.RefreshToken{UserId: uuid.MustParse(userId), StatusLogout: 1})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	return nil
}
