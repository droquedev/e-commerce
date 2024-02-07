package listeners

import (
	"log"

	"github.com/nats-io/stan.go"
)

func InitializeListeners(client stan.Conn) {
	log.Println("Initializing listeners")
	NewUserCreatedListener(client, "products-service").Listen()
}
