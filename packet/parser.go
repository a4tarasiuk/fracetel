package packet

import (
	"bytes"
	"encoding/binary"
	"log"
)

const HeaderTotalBytes = 24

func ParsePacket(rawPacket RawPacket) {
	headerBuffer := bytes.NewBuffer(rawPacket[:HeaderTotalBytes])

	header := Header{}

	err := binary.Read(headerBuffer, binary.LittleEndian, &header)

	if err != nil {
		log.Printf("Error during reading Header: %s", err)
	}

	packetID := ID(header.PacketID)

	if packetID == MotionID || packetID == EventID {
		return
	}

	log.Printf(
		"Packet - [%s]: \"%s\" | %+v\n",
		IDName[ID(header.PacketID)],
		IDDescription[ID(header.PacketID)],
		header,
	)

	if packetID == LapDataID {

		lapData := make([]LapData, 22)

		lapDataBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err = binary.Read(lapDataBuffer, binary.LittleEndian, &lapData)

		if err != nil {
			log.Printf("Error during reading LapData: %s", err)
		}

		for i := 0; i < 22; i++ {
			log.Printf("Lap Data: %+v\n", lapData[i])
		}

	} else if packetID == CarTelemetryID {
		telemetries := make([]CarTelemetry, 22)

		carTelemetryBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err := binary.Read(carTelemetryBuffer, binary.LittleEndian, &telemetries)

		if err != nil {
			log.Printf("Error during reading car telemetries: %s", err)
		}

		for i := 0; i < 22; i++ {
			log.Printf("Car Telemetry: %+v\n", telemetries[i])
		}
	}
}
