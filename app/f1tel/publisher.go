package f1tel

import (
	"fracetel/core/messages"
)

type MessagePublisher interface {
	Publish(message *messages.Message, subject string) error
}
