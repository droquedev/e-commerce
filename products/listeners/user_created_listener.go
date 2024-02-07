package listeners

import (
	"log"

	"github.com/droquedev/e-commerce/pkg/nats"
	"github.com/nats-io/stan.go"
)

type UserCreatedListener struct {
	baseListener *nats.BaseListener
}

func NewUserCreatedListener(client stan.Conn, queGroupName string) nats.Listener {
	baseListener := nats.NewBaseListener(client, "user:created", queGroupName)
	return &UserCreatedListener{
		baseListener: baseListener,
	}
}

func (l *UserCreatedListener) OnMessage(message *stan.Msg) {
	log.Printf("Received user.created event: %s", string(message.Data))
	message.Ack()
}

func (l *UserCreatedListener) Listen() {
	l.baseListener.Listen(l.OnMessage)
}
