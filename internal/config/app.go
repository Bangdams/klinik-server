package config

import (
	"klinikserver/internal/delivery/http"
	"klinikserver/internal/delivery/http/route"
	"klinikserver/internal/repository"
	"klinikserver/internal/usecase/role"
	"klinikserver/internal/usecase/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// repo
	userRepo := repository.NewUserRepository()
	roleRepo := repository.NewRoleRepository()

	// usecase
	userUsecase := user.NewUserUsecase(userRepo, config.DB, config.Validate)
	roleUsecase := role.NewRoleUsecase(roleRepo, config.DB, config.Validate)

	// controller
	userController := http.NewUserController(userUsecase)
	roleController := http.NewRoleController(roleUsecase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		RoleController: roleController,
	}

	routeConfig.Setup()
}
