package controller

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"warrantyapi/configuration"
	"warrantyapi/constant"
	"warrantyapi/model"
	"warrantyapi/service"
)

type PromotionController struct {
	service.PromotionService
	configuration.Config
}

func NewPromotionController(PromotionService *service.PromotionService, config configuration.Config) *PromotionController {
	return &PromotionController{PromotionService: *PromotionService, Config: config}
}

func (controller PromotionController) Route(app *fiber.App) {
	apiV1 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV1.Group("/promotion")
	api.Post("/", controller.create)
	api.Put("/", controller.update)
	api.Get("/", controller.list)
	api.Get("/active", controller.listActive)
}

func (controller PromotionController) create(c *fiber.Ctx) error {
	type promotions struct {
		Promotions []model.PromotionRequest `json:"promotions"`
	}
	promotionsInput := &promotions{}
	if err := c.BodyParser(promotionsInput); err != nil {
		print("An error occurred when parsing the Promotion: " + err.Error())
	}
	for _, promotion := range promotionsInput.Promotions {
		println(promotion.PromotionType)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	result := controller.PromotionService.Create(c.Context(), promotionsInput.Promotions, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

func (controller PromotionController) update(c *fiber.Ctx) error {
	type promotions struct {
		Promotions []model.PromotionRequest `json:"promotions"`
	}
	promotionsInput := &promotions{}
	if err := c.BodyParser(promotionsInput); err != nil {
		print("An error occurred when parsing the Promotion: " + err.Error())
	}
	for _, promotion := range promotionsInput.Promotions {
		println(promotion.ID)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	log.Debug().
		Any("promotions", promotionsInput.Promotions).
		Send()
	result := controller.PromotionService.Update(c.Context(), promotionsInput.Promotions, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller PromotionController) list(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	warrantyId := c.Query("promotion_id")
	log.Debug().
		Int("offset", p.Offset).
		Int("limit", p.Limit).
		Str("order", p.Order).
		Send()
	if p.Limit > 50 {
		p.Limit = 50
	}
	promotionSearch := model.PromotionRequest{
		ID: warrantyId,
	}
	result := controller.PromotionService.List(c.Context(), p.Offset, p.Limit, p.Order, promotionSearch)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}

func (controller PromotionController) listActive(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	promotionType := c.Query("promotion_type")
	promotionBrand := c.Query("promotion_brand")
	warrantyCreated := c.Query("warranty_created")
	warrantyCreatedDate, _ := time.Parse(constant.FORMAT_DATE, warrantyCreated)
	log.Debug().
		Int("offset", p.Offset).
		Int("limit", p.Limit).
		Str("order", p.Order).
		Str("promotionType", promotionType).
		Str("promotionBrand", promotionBrand).
		Str("warrantyCreated", warrantyCreated).
		Time("warrantyCreatedDate", warrantyCreatedDate).
		Send()
	promotionSearch := model.PromotionRequest{
		PromotionStatus: "active",
		PromotionType:   promotionType,
		PromotionBrand:  strings.ToUpper(promotionBrand),
	}

	result := controller.PromotionService.ListActivePromotion(c.Context(), p.Offset, p.Limit, p.Order, promotionSearch, warrantyCreatedDate)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    constant.STATUS_CODE_OK,
		Message: constant.MESSAGE_SUCCESS,
		Data:    result,
	})
}
