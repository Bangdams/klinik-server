package config

import (
	"klinikserver/internal/delivery/http"
	"klinikserver/internal/delivery/http/route"
	"klinikserver/internal/repository"
	"klinikserver/internal/usecase"

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

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepo, config.DB, config.Validate)

	// controller
	userController := http.NewUserController(userUsecase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}
