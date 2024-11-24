package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/messages"
)

type carTelemetryPacketParser struct{}

func (p carTelemetryPacketParser) ToMessage(header *Header, rawPacket RawPacket) (
	*messages.Message,
	error,
) {
	telemetries := make([]carTelemetry, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &telemetries)

	if err != nil {
		log.Printf("Error during reading CarTelemetry: %s", err)
	}

	telemetryPayload := telemetries[header.PlayerCarIdx].ToMessagePayload()

	msg := messages.New(
		messages.CarTelemetryMessageType,
		header.SessionUID,
		header.PacketID,
		header.FrameIdentifier,
		&telemetryPayload,
	)

	log.Printf("Car Telemetry: %+v\n", msg.Payload)

	return &msg, nil
}
