package udp

import (
	"context"
	"log"

	"fracetel/internal/messaging"
)

func messagePublisher(
	eventStream messaging.EventStream,
	telMessageChan <-chan *messaging.Message,
) {
	for telMessage := range telMessageChan {

		topicName, ok := messaging.MessageTypeTopicMap[telMessage.Type]

		if !ok {
			continue
		}

		if err := eventStream.Publish(context.TODO(), topicName, telMessage); err != nil {
			log.Printf("failed to publish message: %+v", telMessage)
		}
	}
}
