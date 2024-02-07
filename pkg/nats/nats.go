package nats

import (
	"os"
	"sync"

	"github.com/nats-io/stan.go"
)

type Nats struct {
	NatsConn stan.Conn
}

var instance *Nats
var once sync.Once

func (n *Nats) Close() error {
	return n.NatsConn.Close()
}

func GetNatsConn() *Nats {
	once.Do(func() {
		nc, err := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"), stan.NatsURL(os.Getenv("NATS_URI")))
		if err != nil {
			panic(err)
		}

		instance = &Nats{
			NatsConn: nc,
		}
	})

	return instance
}
