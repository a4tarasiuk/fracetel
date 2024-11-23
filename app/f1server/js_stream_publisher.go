package f1server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"fracetel/core/messages"
	"github.com/nats-io/nats.go/jetstream"
)

type jetStreamMessagePublisher struct {
	js jetstream.JetStream
}

func NewJetStreamMessagePublisher(js jetstream.JetStream) *jetStreamMessagePublisher {
	return &jetStreamMessagePublisher{js: js}
}

func (p *jetStreamMessagePublisher) Publish(message *messages.Message, subject string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // TODO: 30 sec?
	defer cancel()

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("failed to marshal message before sending to js")
		return err
	}

	_, err = p.js.Publish(ctx, subject, data)

	return err
}
