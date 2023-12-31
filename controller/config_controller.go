package controller

import (
	"github.com/gofiber/fiber/v2"

	"warrantyapi/configuration"
	"warrantyapi/middleware"
	"warrantyapi/model"
	"warrantyapi/service"
)

type ConfigController struct {
	service.ConfigService
	configuration.Config
}

func NewConfigController(ConfigService *service.ConfigService, config configuration.Config) *ConfigController {
	return &ConfigController{ConfigService: *ConfigService, Config: config}
}

func (controller ConfigController) Route(app *fiber.App) {
	apiV1 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV1.Group("/config")
	// api.Post("/", controller.create)
	api.Put("/", middleware.AuthenticateJWT("ROLE_ADMIN"),controller.update)
	api.Get("/", middleware.AuthenticateJWT("ROLE_ADMIN"), controller.list)
}

func (controller ConfigController) update(c *fiber.Ctx) error {
	type configs struct {
		Configs []model.ConfigRequest `json:"configs"`
	}
	configsInput := &configs{}
	// var ectInfos []model.EctInfo
	if err := c.BodyParser(configsInput); err != nil {
		print("An error occurred when parsing the Config: " + err.Error())
	}
	for _, config := range configsInput.Configs {
		// println(config.ID)
		println(config.ConfigCode)
		println(config.ConfigValue)
		// println(config.ConfigDetail)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	result := controller.ConfigService.Update(c.Context(), configsInput.Configs, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller ConfigController) list(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	configCode := c.Query("config_code")
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 50 {
		p.Limit = 50
	}
	configSearch := model.ConfigRequest{
		ConfigCode: configCode,
	}
	result := controller.ConfigService.List(c.Context(), p.Offset, p.Limit, p.Order, configSearch)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}
