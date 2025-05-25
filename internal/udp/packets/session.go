package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
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

func (s session) ToTelemetryMessagePayload() telemetry.Session {
	return telemetry.Session{
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

type SessionParser struct {
}

func (p SessionParser) ToTelemetryMessage(header *Header, rawPacket RawPacket) (
	*messaging.Message,
	error,
) {

	sessionPacket := session{}

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &sessionPacket)

	if err != nil {
		log.Printf("Error during reading Session: %s", err)
	}

	payload := sessionPacket.ToTelemetryMessagePayload()

	msg := messaging.NewMessage(
		telemetry.SessionMessageType,
		header.SessionUID,
		header.FrameIdentifier,
		&payload,
	)

	log.Printf("Session state: %+v\n", msg.Payload)

	return &msg, nil
}
