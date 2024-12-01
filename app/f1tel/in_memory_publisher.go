package f1tel

import "fracetel/core/messages"

type inMemoryMessagePublisher struct {
	messages map[string][]*messages.Message
}

func newInMemoryMessagePublisher() inMemoryMessagePublisher {
	return inMemoryMessagePublisher{
		messages: make(map[string][]*messages.Message),
	}
}

func (p inMemoryMessagePublisher) Publish(message *messages.Message, subject string) error {
	p.messages[subject] = append(p.messages[subject], message)

	return nil
}
