package controller

import (
	"warrantyapi/common"
	"warrantyapi/configuration"
	"warrantyapi/constant"
	"warrantyapi/model"
	"warrantyapi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type AuthenController struct {
	service.AuthenService
	configuration.Config
}

func NewAuthenController(authenService *service.AuthenService, config configuration.Config) *AuthenController {
	return &AuthenController{AuthenService: *authenService, Config: config}
}

func (controller AuthenController) Route(app *fiber.App) {
	api := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
}

// Login get user and pass
func (controller AuthenController) Login(c *fiber.Ctx) error {
	input := new(model.Auth)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	log.Debug().
		Str("userName", input.UserName).
		Str("userPass", input.UserPass).
		Send()
	result := controller.AuthenService.Login(c.Context(), input.UserName, input.UserPass)
	log.Debug().
		Str("userType", result.UserType).
		Send()
	roles := []string{
		"ROLE_USER",
	}
	if result.UserType == "1" {
		roles = []string{
			"ROLE_ADMIN",
			"ROLE_USER",
		}
	}
	var userRoles []map[string]interface{}
	for _, role := range roles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": role,
		})
	}
	tokenJwtResult := common.GenerateToken(result.UserName, userRoles)
	resultWithToken := map[string]interface{}{
		"token":     tokenJwtResult,
		"user_name": result.UserName,
		"roles":     userRoles,
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    resultWithToken,
	})
}
