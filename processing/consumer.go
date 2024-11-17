package processing

import (
	"log"

	"fracetel/models"
)

type messageConsumer struct {
	messageCh <-chan *models.Message

	wsMessageCh chan<- *models.Message
}

func NewMessageConsumer(
	messageCh <-chan *models.Message,
	wsMessageCh chan<- *models.Message,
) *messageConsumer {
	return &messageConsumer{
		messageCh:   messageCh,
		wsMessageCh: wsMessageCh,
	}
}

func (c *messageConsumer) ProcessMessages() {
	for msg := range c.messageCh {
		log.Printf("Consumer recieved a message: %+v\n", msg)
		c.wsMessageCh <- msg
	}
}
