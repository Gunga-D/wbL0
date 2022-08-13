package stream

import (
	"fmt"

	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/nats-io/stan.go"
)

func Init(config *config.NATStreamingConfig, ordStream *OrdersStream) error {
	natsUrl := fmt.Sprintf("nats://%s:%s", config.Host, config.Port)
	stream, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsURL(natsUrl), stan.MaxPubAcksInflight(1000))
	if err != nil {
		return err
	}

	_, err = stream.Subscribe("orders", ordStream.addOrder)
	return err
}
