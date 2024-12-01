package f1tel

import (
	"encoding/json"
	"log"

	"fracetel/core/messages"
	"github.com/nats-io/nats.go"
)

type jsCoreMessagePublisher struct {
	conn *nats.Conn
}

func NewJSCoreMessagePublisher() *jsCoreMessagePublisher {
	conn, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Printf("failed to connect to NATS core: %s", err)
	}

	return &jsCoreMessagePublisher{conn: conn}
}

func (p *jsCoreMessagePublisher) Publish(message *messages.Message, subject string) error {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("failed to marshal message before sending to js")
		return err
	}

	return p.conn.Publish(subject, data)
}
