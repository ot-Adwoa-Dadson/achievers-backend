package services

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to []string, subject, body string) error {
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		port,
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	return d.DialAndSend(m)
}
