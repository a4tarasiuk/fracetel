package packets

import (
	"bytes"
	"encoding/binary"
	"errors"

	"fracetel/core/messages"
)

const HeaderTotalBytes = 24

var PacketByteOrder = binary.LittleEndian

type PacketParser interface {
	ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error)
}

var packetParsersMap = map[ID]PacketParser{
	LapDataID:             lapTimePacketParser{},
	CarTelemetryID:        carTelemetryPacketParser{},
	CarStatusID:           carStatusParser{},
	CarDamageID:           carDamageParser{},
	SessionID:             sessionParser{},
	SessionHistoryID:      sessionHistoryParser{},
	FinalClassificationID: finalClassificationParser{},
}

func GetParserForPacket(packetID ID) (PacketParser, error) {
	parser, exists := packetParsersMap[packetID]

	if !exists {
		return nil, errors.New("packet is not supported")
	}

	return parser, nil
}

func ParserPacketHeader(packet RawPacket) (*Header, error) {
	headerBuffer := bytes.NewBuffer(packet[:HeaderTotalBytes])

	header := Header{}

	err := binary.Read(headerBuffer, PacketByteOrder, &header)

	if err != nil {
		return &Header{}, err
	}

	// log.Printf(
	// 	"Packet - [%s]: \"%s\" | %d\n",
	// 	IDName[ID(header.PacketID)],
	// 	IDDescription[ID(header.PacketID)],
	// 	header.PacketID,
	// )

	return &header, nil
}
