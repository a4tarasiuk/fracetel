package packets

import (
	"log"

	"fracetel/core/models"
)

type lapTimePacketParser struct{}

func (p lapTimePacketParser) ToMessage(header *Header, rawPacket RawPacket) (*models.Message, error) {

	lapData := make([]LapData, F1TotalCars)

	parsePacketBody(rawPacket, &lapData)

	msg := &models.Message{
		Type:      models.LapDataMessageType,
		SessionID: header.SessionUID,
		Payload:   lapData[header.PlayerCarIdx].ToFRT(),
	}

	log.Printf("Lap Data: %+v\n", msg.Payload)

	return msg, nil
}
