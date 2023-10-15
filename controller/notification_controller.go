package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"warrantyapi/configuration"
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

	fmt.Println("Try sending mail...")
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@ppsuperwheels.com")
	m.SetHeader("To", "opas.miracle@gmail.com")
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "ลงทะเบียนรับประกัน")
	m.SetBody("text/html", result)
	// m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("mail.ppsuperwheels.com", 25, "noreply@ppsuperwheels.com", "+PPsuper@1234")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	if err != nil {
		print("Error occurred when send email : " + err.Error())
	}
	return err
}
