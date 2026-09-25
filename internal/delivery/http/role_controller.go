package http

import (
	"klinikserver/internal/model"
	usecase "klinikserver/internal/usecase/role"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	FindByName(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type RoleControllerImpl struct {
	RoleUsecase usecase.RoleUsecase
}

func NewRoleController(roleUsecase usecase.RoleUsecase) RoleController {
	return &RoleControllerImpl{
		RoleUsecase: roleUsecase,
	}
}

// Create implements [RoleController].
func (controller *RoleControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.RoleRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request : ", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.RoleUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create role")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

// ?? cek kemanan na soalna pakai query parameter jang name na kana ai we
// Delete implements [RoleController].
func (controller *RoleControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := controller.RoleUsecase.Delete(ctx.UserContext(), id); err != nil {
		log.Println("failed to delete role")
		return err
	}

	return nil
}

// FindAll implements [RoleController].
func (controller *RoleControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.RoleResponse
	var err error

	name := ctx.Query("name")
	order := ctx.Query("order")
	page := ctx.QueryInt("page")
	limit := ctx.QueryInt("limit")

	responses, currentPage, totalRecords, totalPages, err := controller.RoleUsecase.FindAll(ctx.UserContext(), name, order, page, limit)
	if err != nil {
		log.Println("failed to FindAll role")
		return err
	}

	return ctx.JSON(model.WebResponsesPagination[model.RoleResponse]{
		Data:         responses,
		CurrentPage:  *currentPage,
		TotalRecords: *totalRecords,
		TotalPages:   *totalPages,
	})
}

// FindByName implements [RoleController].
func (controller *RoleControllerImpl) FindByName(ctx *fiber.Ctx) error {
	return nil
}

// Update implements [RoleController].
func (controller *RoleControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateRoleRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("error badrequest:", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.RoleUsecase.Update(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to update role")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}
