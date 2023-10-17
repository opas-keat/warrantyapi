package controller

import (
	"github.com/gofiber/fiber/v2"

	"warrantyapi/configuration"
	"warrantyapi/model"
	"warrantyapi/service"
)

type ProductController struct {
	service.ProductService
	configuration.Config
}

func NewProductController(ProductService *service.ProductService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *ProductService, Config: config}
}

func (controller ProductController) Route(app *fiber.App) {
	apiV1 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV1.Group("/product")
	api.Post("/", controller.create)
	// station.Post("/", middleware.AuthenticateJWT("ROLE_USER"), controller.CreateProduct)
	api.Get("/", controller.list)
}

func (controller ProductController) create(c *fiber.Ctx) error {
	type products struct {
		Products []model.ProductRequest `json:"products"`
	}
	productsInput := &products{}
	if err := c.BodyParser(productsInput); err != nil {
		print("An error occurred when parsing the Product: " + err.Error())
	}
	for _, product := range productsInput.Products {
		println(product.ProductType)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	result := controller.ProductService.Create(c.Context(), productsInput.Products, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller ProductController) list(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	warrantyId := c.Query("warranty_id")
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 50 {
		p.Limit = 50
	}
	productSearch := model.ProductRequest{
		ID: warrantyId,
	}
	result := controller.ProductService.List(c.Context(), p.Offset, p.Limit, p.Order, productSearch)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}
