package email

import (
	"os"

	"github.com/resend/resend-go/v2"
	"gopkg.in/mail.v2"
)

type EmailSender struct {
	client *resend.Client
	dialer *mail.Dialer
}

func NewEmailSender() *EmailSender {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	dialer := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))
	dialer.StartTLSPolicy = mail.MandatoryStartTLS

	return &EmailSender{
		client: client,
		dialer: dialer,
	}
}

func (e *EmailSender) SendEmail(to string, subject string, body string) (*resend.SendEmailResponse, error) {

	params := &resend.SendEmailRequest{
		From:    "No-Reply <no-reply@droque.dev>",
		To:      []string{to},
		Html:    body,
		Subject: subject,
	}

	return e.client.Emails.Send(params)
}

func (e *EmailSender) SendEmailDeprecated(to string, subject string, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return e.dialer.DialAndSend(m)
}
