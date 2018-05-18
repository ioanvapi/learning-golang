package main

import (
	"net/smtp"
	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "Akagi201 <postmaster@mg.ortc.io>"
	e.To = []string{"akagi201@gmail.com"}
	e.Subject = "Hello"
	e.Text = []byte("Testing some Mailgun awesomeness")
	err := e.Send("smtp.mailgun.org:587", smtp.PlainAuth("", "postmaster@mg.ortc.io", "passwd", "smtp.mailgun.org"))
	if err != nil {
		panic(err)
	}
}
