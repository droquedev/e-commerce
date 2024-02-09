package email

import (
	"os"

	"gopkg.in/mail.v2"
)

type EmailSender struct {
	dialer *mail.Dialer
}

func NewEmailSender() *EmailSender {

	dialer := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))
	dialer.StartTLSPolicy = mail.MandatoryStartTLS

	return &EmailSender{
		dialer: dialer,
	}
}

func (e *EmailSender) SendEmail(to string, subject string, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return e.dialer.DialAndSend(m)
}
