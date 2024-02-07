package nats

import (
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type Listener interface {
	Listen()
	OnMessage(message *stan.Msg)
}

type BaseListener struct {
	subject        string
	queueGroupName string
	client         stan.Conn
	ackWait        time.Duration
}

func NewBaseListener(client stan.Conn, subject string, queueGroupName string) *BaseListener {
	return &BaseListener{
		subject:        subject,
		queueGroupName: queueGroupName,
		client:         client,
		ackWait:        5 * time.Second,
	}
}

func (l *BaseListener) Listen(handler stan.MsgHandler) {
	_, err := l.client.QueueSubscribe(l.subject,
		l.queueGroupName,
		handler,
		stan.SetManualAckMode(),
		stan.DeliverAllAvailable(),
		stan.AckWait(l.ackWait),
		stan.DurableName(l.queueGroupName),
	)

	if err != nil {
		log.Printf("Error subscribing to %s: %v\n", l.subject, err)
	}
}
