package listeners

import (
	"log"

	"github.com/droquedev/e-commerce/notification-service/pkg/email"
	"github.com/nats-io/stan.go"
)

func InitializeListeners(client stan.Conn, emailSender *email.EmailSender) {
	log.Println("Initializing listeners")
	NewUserCreatedListener(client, emailSender, "notification-service").Listen()
}
