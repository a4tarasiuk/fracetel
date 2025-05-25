package packets

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
)

type finalClassification struct {
	FinishingPosition uint8

	TotalLaps uint8

	StartingPosition uint8

	Points uint8

	NumPitStops uint8

	ResultStatus uint8

	BestLapTimeMs uint32

	TotalRaceTimeSec float64
	PenaltiesTimeSec uint8

	TotalPenalties uint8

	TotalTyreStints uint8

	TyreStintsActual  [8]uint8
	TyreStintsVisual  [8]uint8
	TyreStintsEndLaps [8]uint8
}

func (fc finalClassification) ToTelemetryMessagePayload(header *Header) telemetry.FinalClassification {
	return telemetry.FinalClassification{
		SessionID:         header.SessionUID,
		FrameIdentifier:   header.FrameIdentifier,
		OccurredAt:        time.Now().UTC(),
		FinishingPosition: int(fc.FinishingPosition),
		StartingPosition:  int(fc.StartingPosition),
		BestLapTimeMs:     float32(fc.BestLapTimeMs),
	}
}

type FinalClassificationParser struct{}

func (p FinalClassificationParser) ToTelemetryMessage(header *Header, rawPacket RawPacket) (
	*messaging.Message,
	error,
) {

	finalClassificationPackets := make([]finalClassification, F1TotalCars)

	// 1 - is uint8 that is responsible for number of cars
	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes+1:])

	err := binary.Read(buffer, PacketByteOrder, &finalClassificationPackets)

	if err != nil {
		log.Printf("Error during reading Session: %s", err)
	}

	payload := finalClassificationPackets[header.PlayerCarIdx].ToTelemetryMessagePayload(header)

	msg := messaging.NewMessage(
		messaging.FinalClassificationMessageType,
		header.SessionUID,
		header.FrameIdentifier,
		&payload,
	)

	log.Printf("Final classification: %+v\n", msg.Payload)

	return &msg, nil
}
