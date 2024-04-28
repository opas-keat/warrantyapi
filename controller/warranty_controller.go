package controller

import (
	"warrantyapi/configuration"
	"warrantyapi/middleware"
	"warrantyapi/model"
	"warrantyapi/service"

	"github.com/gofiber/fiber/v2"
)

type WarrantyController struct {
	service.WarrantyService
	configuration.Config
}

func NewWarrantyController(warrantyService *service.WarrantyService, config configuration.Config) *WarrantyController {
	return &WarrantyController{WarrantyService: *warrantyService, Config: config}
}

func (controller WarrantyController) Route(app *fiber.App) {
	apiV2 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV2.Group("/warranty")
	// api.Post("/", middleware.AuthenticateJWT("ROLE_USER"), controller.create)
	// api.Put("/", middleware.AuthenticateJWT("ROLE_USER"), controller.update)
	// api.Delete("/:id", middleware.AuthenticateJWT("ROLE_USER"), controller.delete)
	api.Post("/", controller.create)
	api.Put("/", controller.update)
	api.Delete("/:id", controller.delete)
	api.Get("/customer", controller.listCustomer)
	api.Get("/excels", controller.listExcels)
	api.Get("/", controller.list)
	api.Get("/:id", controller.findById)
	api.Static("/uploads", "./uploads")
}

func (controller WarrantyController) create(c *fiber.Ctx) error {
	type warranty struct {
		Warranty model.WarrantyRequest `json:"warranty"`
	}
	warrantysInput := &warranty{}
	if err := c.BodyParser(warrantysInput); err != nil {
		print("An error occurred when parsing the warrantys: " + err.Error())
	}
	println(warrantysInput.Warranty.WarrantyNo)
	println(warrantysInput.Warranty.DealerCode)
	println(warrantysInput.Warranty.CustomerName)
	for _, product := range warrantysInput.Warranty.ProductRequest {
		println(product.ProductBrand)
	}
	// userName := middleware.GetUserNameFromToken(c)
	userName := ""
	result := controller.WarrantyService.Create(c.Context(), warrantysInput.Warranty, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) update(c *fiber.Ctx) error {
	type warrantys struct {
		Warrantys []model.WarrantyRequest `json:"warrantys"`
	}
	warrantysInput := &warrantys{}
	if err := c.BodyParser(warrantysInput); err != nil {
		print("An error occurred when parsing the Warranty: " + err.Error())
	}
	for _, warranty := range warrantysInput.Warrantys {
		println(warranty.ID)
	}
	userName := middleware.GetUserNameFromToken(c)
	result := controller.WarrantyService.Update(c.Context(), warrantysInput.Warrantys, userName)
	// for _, rs := range result {
	// 	println(rs.WarrantyNo)
	// 	println(rs.ID)
	// 	path := common.CreatePathFileForUpload(rs.WarrantyNo, rs.ID)
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		println(err)
	// 	}
	// 	println(file.Filename)
	// 	c.SaveFile(file, fmt.Sprintf(path+"%s", rs.ID+"_1.png"))
	// }
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	println(id)
	userName := middleware.GetUserNameFromToken(c)
	result := controller.WarrantyService.Delete(c.Context(), id, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) listCustomer(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	customerPhone := c.Query("customer_phone")
	CustomerLicensePlate := c.Query("customer_license_plate")
	CustomerEmail := c.Query("customer_email")
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 50 {
		p.Limit = 50
	}
	warrantyInput := model.WarrantyRequest{
		CustomerPhone:        customerPhone,
		CustomerLicensePlate: CustomerLicensePlate,
		CustomerEmail:        CustomerEmail,
	}
	result := controller.WarrantyService.ListCustomer(c.Context(), p.Offset, p.Limit, p.Order, warrantyInput)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) list(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	// warrantyFirstName := c.Query("warranty_first_name")
	// warrantySurName := c.Query("warranty_sur_name")
	// warrantyAgency := c.Query("warranty_agency")
	// warrantyAffiliate := c.Query("warranty_affiliate")
	// warrantyTelePhone := c.Query("warranty_telephone")
	customerPhone := c.Query("customer_phone")
	CustomerLicensePlate := c.Query("customer_license_plate")
	CustomerEmail := c.Query("customer_email")
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 50 {
		p.Limit = 50
	}
	warrantyInput := model.WarrantyRequest{
		CustomerPhone:        customerPhone,
		CustomerLicensePlate: CustomerLicensePlate,
		CustomerEmail:        CustomerEmail,
	}
	result := controller.WarrantyService.List(c.Context(), p.Offset, p.Limit, p.Order, warrantyInput)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) findById(c *fiber.Ctx) error {
	id := c.Params("id")
	println(id)
	result := controller.WarrantyService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    result,
	})
}

func (controller WarrantyController) listExcels(c *fiber.Ctx) error {
	p := new(model.ListReq)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	// customerPhone := c.Query("customer_phone")
	CustomerLicensePlate := c.Query("customer_license_plate")
	// CustomerEmail := c.Query("customer_email")
	println(p.Offset)
	println(p.Limit)
	println(p.Order)
	if p.Limit > 5 {
		p.Limit = 5
	}
	warrantyInput := model.WarrantyRequest{
		// CustomerPhone:        customerPhone,
		CustomerLicensePlate: CustomerLicensePlate,
		// CustomerEmail:        CustomerEmail,
	}
	result := controller.WarrantyService.ListExcels(c.Context(), p.Offset, p.Limit, p.Order, warrantyInput)
	return c.SendFile(result)
}
