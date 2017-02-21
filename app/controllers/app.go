package controllers

import (
	"crypto/tls"
	"os"

	"github.com/grantmd/go-coinbase"
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

func (c App) Trader() revel.Result {
	apiKey := os.Getenv("COINBASE_API_KEY")

	if apiKey != "" {
		// load all data for the previous n months
		c := coinbase.Client{
			APIKey: apiKey,
		}

		historical, err := c.GetHistoricalPrices(1)
		if err != nil {
			revel.WARN.Println("Error getting historical prices:", err.Error())
		}

		dir, _ := os.Getwd()

		dataOut, err := os.OpenFile(dir+"/public/data/data.csv", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			revel.WARN.Println("Error opening data file:", err.Error())
		}

		dataOut.WriteString("date,price\n")
		dataOut.WriteString(historical)

		dataOut.Close()
	} else {
		revel.WARN.Println("Missing \"COINBASE_API_KEY\" env variable")
	}

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

		email := os.Getenv("EMAIL_USERNAME")
		emailPassword := os.Getenv("EMAIL_PASSWORD")

		d := gomail.NewDialer("smtp.gmail.com", 587, email, emailPassword)
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
