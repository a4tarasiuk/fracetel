package parser

import (
	"bytes"
	"encoding/binary"
	"log"

	packets2 "fracetel/f1server/packets"
)

const HeaderTotalBytes = 24

func ParsePacket(rawPacket packets2.RawPacket) {
	headerBuffer := bytes.NewBuffer(rawPacket[:HeaderTotalBytes])

	header := packets2.Header{}

	err := binary.Read(headerBuffer, binary.LittleEndian, &header)

	if err != nil {
		log.Printf("Error during reading Header: %s", err)
	}

	packetID := packets2.ID(header.PacketID)

	if packetID == packets2.MotionID || packetID == packets2.EventID {
		return
	}

	if packetID == packets2.LapDataID {

		lapData := make([]packets2.LapData, packets2.F1TotalCars)

		lapDataBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err = binary.Read(lapDataBuffer, binary.LittleEndian, &lapData)

		if err != nil {
			log.Printf("Error during reading LapData: %s", err)
		}

		playerLapData := lapData[header.PlayerCarIdx].ToFRT()

		log.Printf("Lap Data: %+v\n", playerLapData)

	} else if packetID == packets2.CarTelemetryID {
		telemetries := make([]packets2.CarTelemetry, packets2.F1TotalCars)

		carTelemetryBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

		err = binary.Read(carTelemetryBuffer, binary.LittleEndian, &telemetries)

		if err != nil {
			log.Printf("Error during reading car telemetries: %s", err)
		}

		playerCarTelemetry := telemetries[header.PlayerCarIdx]

		frtCarTelemetry := playerCarTelemetry.ToFRT()

		log.Printf("Car Telemetry: %+v\n", frtCarTelemetry)

	} else {
		log.Printf(
			"Packet - [%s]: \"%s\" | %+v\n",
			packets2.IDName[packets2.ID(header.PacketID)],
			packets2.IDDescription[packets2.ID(header.PacketID)],
			header,
		)
	}
}
