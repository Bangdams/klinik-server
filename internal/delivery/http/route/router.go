package route

import (
	"klinikserver/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController http.UserController
}

func (config *RouteConfig) Setup() {
	// Api for login
	config.App.Post("/login", config.UserController.Login)
	config.App.Post("/logout", config.UserController.Logout)

	// Group Api
	api := config.App.Group("/api")

	// Api For Management User
	api.Get("/users", config.UserController.FindAll)
	// api.Get("/users/:id", util.CheckLevel("admin"), config.UserController.FindByIdForUpdate)
	// api.Post("/users", util.CheckLevel("admin"), config.UserController.Create)
	// api.Delete("/users/:id", util.CheckLevel("admin"), config.UserController.Delete)
	// api.Put("/users", util.CheckLevel("admin"), config.UserController.Update)
}
