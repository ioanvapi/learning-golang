package main

import (
	"net/smtp"
	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "Akagi201 <admin201@blockgw.com>"
	e.To = []string{"akagi201@gmail.com"}
	e.Subject = "Hello"
	e.Text = []byte("Testing some Mailgun awesomeness")
	err := e.Send("smtpcloud.sohu.com:25", smtp.PlainAuth("", "admin201", "DNjMEC79v5rowaAh", "smtpcloud.sohu.com"))
	if err != nil {
		panic(err)
	}
}
