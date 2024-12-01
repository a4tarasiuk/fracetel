package f1tel

import (
	"context"
	"log"

	"fracetel/core/telemetry"
	"fracetel/internal/messaging"
)

func TelemetryMessageProcessor(
	eventStream messaging.EventStream,
	telMessageChan <-chan *telemetry.Message,
) {
	for telMessage := range telMessageChan {

		if err := eventStream.Publish(context.TODO(), telemetry.FRaceTelTopicName, telMessage); err != nil {
			log.Printf("failed to publish message: %+v", telMessage)
		}
	}
}
