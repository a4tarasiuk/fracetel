package f1server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"

	"fracetel/f1server/packets"
	"fracetel/models"
)

const HeaderTotalBytes = 24

func ParsePacket(rawPacket packets.RawPacket) (*models.Message, error) {
	headerBuffer := bytes.NewBuffer(rawPacket[:HeaderTotalBytes])

	header := packets.Header{}

	err := binary.Read(headerBuffer, binary.LittleEndian, &header)

	if err != nil {
		log.Printf("Error during reading Header: %s", err)
	}

	packetID := packets.ID(header.PacketID)

	if packetID == packets.MotionID || packetID == packets.EventID {
		return nil, errors.New("packet was ignored")
	}

	if packetID == packets.LapDataID {

		lapData := make([]packets.LapData, packets.F1TotalCars)

		lapDataBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err = binary.Read(lapDataBuffer, binary.LittleEndian, &lapData)

		if err != nil {
			log.Printf("Error during reading LapData: %s", err)
		}

		msg := &models.Message{
			Type:      models.LapDataMessageType,
			SessionID: header.SessionUID,
			Payload:   lapData[header.PlayerCarIdx].ToFRT(),
		}

		log.Printf("Lap Data: %+v\n", msg.Payload)

		return msg, nil

	} else if packetID == packets.CarTelemetryID {
		telemetries := make([]packets.CarTelemetry, packets.F1TotalCars)

		carTelemetryBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err = binary.Read(carTelemetryBuffer, binary.LittleEndian, &telemetries)

		if err != nil {
			log.Printf("Error during reading car telemetries: %s", err)
		}

		msg := &models.Message{
			Type:      models.CarTelemetryMessageType,
			SessionID: header.SessionUID,
			Payload:   telemetries[header.PlayerCarIdx].ToFRT(),
		}

		log.Printf("Car Telemetry: %+v\n", msg.Payload)

		return msg, nil

	} else {
		log.Printf(
			"Packet - [%s]: \"%s\" | %+v\n",
			packets.IDName[packets.ID(header.PacketID)],
			packets.IDDescription[packets.ID(header.PacketID)],
			header,
		)
	}

	return nil, nil
}
