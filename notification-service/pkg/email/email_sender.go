package email

import (
	"os"

	"github.com/resend/resend-go/v2"
)

type EmailSender struct {
	client *resend.Client
}

func NewEmailSender() *EmailSender {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	return &EmailSender{
		client: client,
	}
}

func (e *EmailSender) SendEmail(to string, subject string, body string) (*resend.SendEmailResponse, error) {

	params := &resend.SendEmailRequest{
		From:    "No-Reply <no-reply@droque.dev>",
		To:      []string{to},
		Html:    body,
		Subject: subject,
		Headers: map[string]string{
			"X-Entity-Ref-ID": "123456789",
		},
	}

	return e.client.Emails.Send(params)
}

/* package email

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
*/
