package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"warrantyapi/configuration"
	"warrantyapi/middleware"
	"warrantyapi/model"
	"warrantyapi/service"
)

type DealerController struct {
	service.DealerService
	configuration.Config
}

func NewDealerController(DealerService *service.DealerService, config configuration.Config) *DealerController {
	return &DealerController{DealerService: *DealerService, Config: config}
}

func (controller DealerController) Route(app *fiber.App) {
	apiV1 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV1.Group("/dealer")
	api.Post("/", middleware.AuthenticateJWT("ROLE_USER"), controller.create)
	// station.Post("/", middleware.AuthenticateJWT("ROLE_USER"), controller.CreateDealer)
	api.Get("/", controller.list)
	// api.Get("/:id", controller.findById)
}

func (controller DealerController) create(c *fiber.Ctx) error {
	type dealers struct {
		Dealers []model.DealerRequest `json:"dealers"`
	}
	dealersInput := &dealers{}
	if err := c.BodyParser(dealersInput); err != nil {
		print("An error occurred when parsing the Dealer: " + err.Error())
	}
	for _, dealer := range dealersInput.Dealers {
		println(dealer.DealerCode)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	result := controller.DealerService.Create(c.Context(), dealersInput.Dealers, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller DealerController) list(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	dealerCode := c.Query("dealer_code")
	dealerId := c.Query("dealer_id")
	println("dealerCode: " + dealerCode)
	println("dealerId: " + dealerId)
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 50 {
		p.Limit = 50
	}
	id, _ := uuid.Parse(dealerId)
	var result []model.DealerResponse
	println("id: " + id.String())
	dealerSearch := model.DealerRequest{
		ID:         id.String(),
		DealerCode: dealerCode,
	}
	if dealerCode != "" || id.String() != "00000000-0000-0000-0000-000000000000" {
		result = controller.DealerService.List(c.Context(), p.Offset, p.Limit, p.Order, dealerSearch)
	} else {
		result = []model.DealerResponse{}
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

// func (controller DealerController) findById(c *fiber.Ctx) error {
// 	id := c.Query("id")
// 	println(id)
// 	result := controller.DealerService.FindById(c.Context(), id)
// 	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
// 		Code:    "000",
// 		Message: "Success",
// 		Data:    result,
// 	})
// }
