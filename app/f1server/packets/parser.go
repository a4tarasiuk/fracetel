package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"

	"fracetel/core/models"
)

const HeaderTotalBytes = 24

type PacketParser interface {
	ToMessage(header *Header, rawPacket RawPacket) (*models.Message, error)
}

var packetParsersMap = map[ID]PacketParser{
	EventID:        eventPacketParser{},
	LapDataID:      lapTimePacketParser{},
	CarTelemetryID: carTelemetryPacketParser{},
}

func GetParserForPacket(packetID ID) (PacketParser, error) {
	parser, exists := packetParsersMap[packetID]

	if !exists {
		return nil, errors.New("packet is not supported")
	}

	return parser, nil
}

func parsePacketBody(packet RawPacket, target interface{}) {
	buffer := bytes.NewBuffer(packet[HeaderTotalBytes:])

	err := binary.Read(buffer, binary.LittleEndian, &target)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}
}

func ParserPacketHeader(packet RawPacket) (*Header, error) {
	headerBuffer := bytes.NewBuffer(packet[:HeaderTotalBytes])

	header := Header{}

	err := binary.Read(headerBuffer, binary.LittleEndian, &header)

	if err != nil {
		return &Header{}, err
	}

	// log.Printf(
	// 	"Packet - [%s]: \"%s\" | %+v\n",
	// 	packets.IDName[packets.ID(header.PacketID)],
	// 	packets.IDDescription[packets.ID(header.PacketID)],
	// 	header,
	// )

	return &header, nil
}
