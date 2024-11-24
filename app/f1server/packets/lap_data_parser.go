package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/messages"
)

type lapTimePacketParser struct{}

func (p lapTimePacketParser) ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error) {

	lapDataPacket := make([]lapData, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &lapDataPacket)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	lapData := lapDataPacket[header.PlayerCarIdx].ToMessagePayload()

	msg := messages.New(
		messages.LapDataMessageType,
		header.SessionUID,
		header.PacketID,
		header.FrameIdentifier,
		&lapData,
	)

	log.Printf("Lap Data: %+v\n", msg.Payload)

	return &msg, nil
}
