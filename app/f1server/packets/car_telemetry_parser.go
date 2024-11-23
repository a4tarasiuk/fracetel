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
	telemetries := make([]CarTelemetry, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, binary.LittleEndian, &telemetries)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	frtTelemetry := telemetries[header.PlayerCarIdx].ToFRT()

	msg := messages.New(messages.CarTelemetryMessageType, header.SessionUID, &frtTelemetry)

	log.Printf("Car Telemetry: %+v\n", msg.Payload)

	return &msg, nil
}
