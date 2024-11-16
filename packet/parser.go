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

	binary.Read(headerBuffer, binary.LittleEndian, &header)

	if header.PacketID == LapDataID {
		// log.Printf("Header: %+v\n", header)

		lapData := LapData{}

		lapDataBuffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes+1:])

		binary.Read(lapDataBuffer, binary.LittleEndian, &lapData)

		log.Printf("Lap Data: %+v\n", lapData)
	}
}
