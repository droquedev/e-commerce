package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/droquedev/e-commerce/notification-service/internal/listeners"
	"github.com/droquedev/e-commerce/notification-service/pkg/email"
	"github.com/droquedev/e-commerce/pkg/nats"
)

func main() {
	log.Println("Notification service started")

	natsConn := nats.GetNatsConn().NatsConn
	emailSender := email.NewEmailSender()
	listeners.InitializeListeners(natsConn, emailSender)

	sigChann := make(chan os.Signal, 1)
	signal.Notify(sigChann, syscall.SIGINT, syscall.SIGTERM)
	<-sigChann

	log.Println("Notification service stopped")
}
