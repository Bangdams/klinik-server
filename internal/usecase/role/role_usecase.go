package role

import (
	"context"
	"klinikserver/internal/model"
	"klinikserver/internal/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RoleUsecase interface {
	FindByName(ctx context.Context, name string) (*model.RoleResponse, error)
	FindAll(ctx context.Context, name string, order string, page int, limit int) (*[]model.RoleResponse, *int, *int, *int, error)
	Create(ctx context.Context, request *model.RoleRequest) (*model.RoleResponse, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, request *model.UpdateRoleRequest) (*model.RoleResponse, error)
}

type RoleUsecaseImpl struct {
	RoleRepo repository.RoleRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewRoleUsecase(roleRepo repository.RoleRepository, DB *gorm.DB, validate *validator.Validate) RoleUsecase {
	return &RoleUsecaseImpl{
		RoleRepo: roleRepo,
		DB:       DB,
		Validate: validate,
	}
}
