package controller

import (
	"strconv"
	"warrantyapi/configuration"
	"warrantyapi/constant"
	"warrantyapi/middleware"
	"warrantyapi/model"
	"warrantyapi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	service.UserService
	configuration.Config
}

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

func (controller UserController) Route(app *fiber.App) {
	api := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	users := api.Group("/users")
	// users.Post("/", controller.create)
	users.Post("/", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.create)
	users.Put("/", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.Update)
	users.Delete("/", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.Delete)
	users.Get("/", middleware.AuthenticateJWT("ROLE_USER"), controller.GetUserByToken)
	users.Get("/search", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.FindAll)
	users.Get("/:id", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.FindById)
}

// create
func (controller UserController) create(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userName := middleware.GetUserNameFromToken(c)
	result := controller.UserService.Create(c.Context(), request, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

// Update
func (controller UserController) Update(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userName := middleware.GetUserNameFromToken(c)
	result := controller.UserService.Update(c.Context(), request, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

// Delete user
func (controller UserController) Delete(c *fiber.Ctx) error {
	var request model.Delete
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("userId", strconv.Itoa(request.ID)).Send()
	userName := middleware.GetUserNameFromToken(c)
	controller.UserService.Delete(c.Context(), strconv.Itoa(request.ID), userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    "",
	})
}

// Get User By Token
func (controller UserController) GetUserByToken(c *fiber.Ctx) error {
	userName := middleware.GetUserNameFromToken(c)
	log.Debug().Str("username", userName).Send()
	result := controller.UserService.FindByUserName(c.Context(), userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

// FindAll
func (controller UserController) FindAll(c *fiber.Ctx) error {
	// log.Debug().Str("FindAll", "").Send()
	limit := c.QueryInt("limit", 20)
	page := c.QueryInt("page", 0)
	orderBy := c.Query("order_by", "")
	sort := c.Query("sort", "")
	// log.Debug().Int("Limit", limit).Send()
	// log.Debug().Int("Offset", page).Send()
	result := controller.UserService.FindAll(c.Context(), limit, page, orderBy, sort)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

// Get User By Id
func (controller UserController) FindById(c *fiber.Ctx) error {
	userId := c.Params("id")
	log.Debug().Str("userId", userId).Send()
	result := controller.UserService.FindById(c.Context(), userId)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}
