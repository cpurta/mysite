package controllers

import (
	"crypto/tls"

	"github.com/revel/revel"
	gomail "gopkg.in/gomail.v2"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Powerlifting() revel.Result {
	return c.Render()
}

func (c App) Quines() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Projects() revel.Result {
	return c.Render()
}

func (c App) Encryption() revel.Result {
	return c.Render()
}

func (c App) Contact(name, email, phone, message string) revel.Result {
	if c.Request.Method == "POST" {
		c.Validation.Required(name).Message("Your name is required")
		c.Validation.Required(email).Message("Your email is required")
		c.Validation.MinSize(message, 20).Message("Your message must be at least 20 characters long.")

		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()

			return c.Render()
		}

		d := gomail.NewDialer("smtp.gmail.com", 587, "cpurta@gmail.com", "rwjavdofhsrkpxqg")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		m := gomail.NewMessage()
		m.SetHeader("From", email)
		m.SetHeader("To", "cpurta@gmail.com")
		m.SetHeader("Subject", "New contact form from: "+name)
		m.SetBody("text/html", message)

		if err := d.DialAndSend(m); err != nil {
			revel.ERROR.Println("Error sending email:", err.Error())
		}
	}

	return c.Render()
}
