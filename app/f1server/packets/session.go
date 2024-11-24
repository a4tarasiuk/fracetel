package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/messages"
)

type session struct {
	Weather uint8

	TrackTemperature int8
	AirTemperature   int8

	TotalLaps   uint8
	TrackLength uint16
	TrackID     int8

	Type uint8

	TimeLeft uint16
	Duration uint16
}

func (s session) ToMessagePayload() messages.Session {
	return messages.Session{
		Weather:          int(s.Weather),
		TrackTemperature: int(s.TrackTemperature),
		AirTemperature:   int(s.AirTemperature),
		TotalLaps:        int(s.TotalLaps),
		TrackLength:      int(s.TrackLength),
		TrackID:          int(s.TrackID),
		Type:             int(s.Type),
		TimeLeft:         int(s.TimeLeft),
		Duration:         int(s.Duration),
	}
}

type sessionParser struct {
}

func (p sessionParser) ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error) {

	sessionPacket := session{}

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &sessionPacket)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	payload := sessionPacket.ToMessagePayload()

	msg := messages.New(
		messages.SessionMessageType,
		header.SessionUID,
		header.PacketID,
		header.FrameIdentifier,
		&payload,
	)

	log.Printf("Session state: %+v\n", msg.Payload)

	return &msg, nil
}
