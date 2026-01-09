package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"warrantyapi/configuration"
	"warrantyapi/model"
	"warrantyapi/service"

	"github.com/gofiber/fiber/v2"

	gomail "gopkg.in/gomail.v2"
)

type NotificationController struct {
	service.WarrantyService
	configuration.Config
}

func NewNotificationController(warrantyService *service.WarrantyService, config configuration.Config) *NotificationController {
	return &NotificationController{WarrantyService: *warrantyService, Config: config}
}

func (controller NotificationController) Route(app *fiber.App) {
	apiV2 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV2.Group("/notification")
	api.Post("/email", controller.email)
}

func (controller NotificationController) email(c *fiber.Ctx) error {
	type w struct {
		Id string `json:"id"`
	}
	warrantysInput := &w{}
	if err := c.BodyParser(warrantysInput); err != nil {
		print("An error occurred when parsing the warrantys: " + err.Error())
	}
	fmt.Println("warrantysInput.Id : ", warrantysInput.Id)
	warranty := controller.WarrantyService.FindById(c.Context(), warrantysInput.Id)

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("./templates/template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, warranty); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	fmt.Println("warranty : ", warranty)
	fmt.Println("Try sending mail...", warranty.CustomerEmail)
	m := gomail.NewMessage()
	m.SetHeader("From", "warranty@ppsuperwheels.com")
	// m.SetHeader("To", warranty.CustomerEmail)
	m.SetHeader("To", "opas.miracle@gmail.com")
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "ลงทะเบียนรับประกัน")
	m.SetBody("text/html", result)
	m.Embed("./templates/logo.png")
	m.Embed("./templates/lineid.png")
	// m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("mail.ppsuperwheels.com", 587, "warranty@ppsuperwheels.com", "+PPsuper@1234")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	if err != nil {
		fmt.Println("Error:", err)
		// print("Error occurred when send email : " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    "",
	})
}
