package udp

import (
	"context"
	"log"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
)

func messagePublisher(
	eventStream messaging.EventStream,
	telMessageChan <-chan *telemetry.Message,
) {
	for telMessage := range telMessageChan {

		topicName, ok := telemetry.MessageTypeTopicMap[telMessage.Type]

		if !ok {
			continue
		}

		if err := eventStream.Publish(context.TODO(), topicName, telMessage); err != nil {
			log.Printf("failed to publish message: %+v", telMessage)
		}
	}
}
