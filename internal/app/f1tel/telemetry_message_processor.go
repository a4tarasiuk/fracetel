package f1tel

import (
	"context"
	"log"

	"fracetel/internal/core/telemetry"
	"fracetel/internal/messaging"
)

func TelemetryMessageProcessor(
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
