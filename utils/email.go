package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

// SendEmail sends an email using Mailtrap SMTP
func SendEmail(to, subject, body string) error {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
