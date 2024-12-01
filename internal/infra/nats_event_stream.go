package infra

import (
	"context"
	"encoding/json"
	"log"

	"fracetel/internal/messaging"
	"github.com/nats-io/nats.go"
)

type natsEventStream struct {
	conn *nats.Conn
}

func NewNatsEventStream(conn *nats.Conn) *natsEventStream {
	return &natsEventStream{conn: conn}
}

func (p *natsEventStream) Publish(ctx context.Context, topicName string, value messaging.Event) error {
	data, err := json.Marshal(value)

	if err != nil {
		log.Printf("failed to marshal message before sending to js")
		return err
	}

	return p.conn.Publish(topicName, data)
}
