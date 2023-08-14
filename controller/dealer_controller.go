package controller

import (
	"github.com/gofiber/fiber/v2"

	"warrantyapi/configuration"
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
	api := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	dealer := api.Group("/dealer")
	dealer.Post("/", controller.CreateDealer)
	// station.Post("/", middleware.AuthenticateJWT("ROLE_USER"), controller.CreateDealer)
	// station.Get("/", controller.ListDealer)
	// station.Get("/:id", controller.FindById)
}

func (controller DealerController) CreateDealer(c *fiber.Ctx) error {
	// type Dealers struct {
	// 	Dealers model.DealerRequest `json:"dealer"`
	// }
	dealerInput := new(model.DealerRequest)
	if err := c.BodyParser(dealerInput); err != nil {
		print("An error occurred when parsing the Dealer: " + err.Error())
	}
	print(dealerInput.DealerCode)
	// userName := middleware.GetUserNameFromToken(c)
	userName := "test"
	result := controller.DealerService.Create(c.Context(), *dealerInput, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}
