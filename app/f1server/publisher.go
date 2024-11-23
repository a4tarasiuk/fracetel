package f1server

import (
	"fracetel/core/messages"
)

type MessagePublisher interface {
	Publish(message *messages.Message, subject string) error
}
