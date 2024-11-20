package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/messages"
)

type lapTimePacketParser struct{}

func (p lapTimePacketParser) ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error) {

	lapData := make([]LapData, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, binary.LittleEndian, &lapData)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	frtLapData := lapData[header.PlayerCarIdx].ToFRT()

	msg := messages.New(messages.LapDataMessageType, header.SessionUID, &frtLapData)

	log.Printf("Lap Data: %+v\n", msg.Payload)

	return &msg, nil
}
