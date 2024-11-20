package packets

import (
	"log"

	"fracetel/core/models"
)

type carTelemetryPacketParser struct{}

func (p carTelemetryPacketParser) ToMessage(header *Header, rawPacket RawPacket) (
	*models.Message,
	error,
) {
	telemetries := make([]CarTelemetry, F1TotalCars)

	parsePacketBody(rawPacket, &telemetries)

	msg := &models.Message{
		Type:      models.CarTelemetryMessageType,
		SessionID: header.SessionUID,
		Payload:   telemetries[header.PlayerCarIdx].ToFRT(),
	}

	log.Printf("Car Telemetry: %+v\n", msg.Payload)

	return msg, nil
}
