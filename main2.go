package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type info struct {
	Name string
}

func (i info) sendMail() {

	// var results model.WarrantyResponse
	// results = model.WarrantyResponse{
	// 	WarrantyNo: "1111",
	// }

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("./templates/template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
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

	// d := gomail.NewDialer("mail.ppsuperwheels.com", 25, "noreply@ppsuperwheels.com", "+PPsuper@1234")
	// m := gomail.NewMessage()
	// m.SetHeader("From", "noreply@ppsuperwheels.com")
	// m.SetHeader("To", "opas.miracle@gmail.com")
	// m.SetHeader("Subject", "ลงทะเบียนรับประกัน")
	// // m.SetBody("text/html", "This is a test")
	// m.SetBody("text/html", result)
	// // m.Attach("template.html")// attach whatever you want
	// if err := d.DialAndSend(m); err != nil {
	// 	fmt.Println("Failed sending mail")
	// 	panic(err)
	// }
	// fmt.Println("Mail sent without error")
}

func main2() {

	d := info{"jack"}

	d.sendMail()
	
}
