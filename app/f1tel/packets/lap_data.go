package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/messages"
)

type lapData struct {
	LastLapTimeMs    uint32
	CurrentLapTimeMs uint32

	FirstSectorTimeMs  uint16
	SecondSectorTimeMs uint16

	LapDistance   float32
	TotalDistance float32

	SafetyCarDelta float32

	CarPosition uint8

	CurrentLapNum uint8

	PitStatus   uint8 // 0 - none, 1 - pitting, 2 - in pit area
	NumPitStops uint8

	Sector uint8 // 0 - sector1, 1 - sector2, 2 - sector3

	CurrentLapInvalid uint8 // 0 - valid, 1 - invalid

	PenaltiesSeconds uint8 //

	Warnings uint8 // total warnings

	NumUnservedDriveThroughPens uint8
	NumUnservedStopsGoPens      uint8

	StartingGridPosition uint8

	DriverStatus uint8 // 0 - in garage, 1 - flying lap, 2 - in lap, 3 - out lap, 4 - on track

	// 0 - invalid, 1 - inactive, 2 - active, 3 - finished, 4 - didnotfinish, 5 - disqualified, 6 - not classified
	// 7 - retired
	ResultStatus uint8

	PitlaneTimerActive  uint8 // 0 - inactive, 1 - active
	PitlaneTimeInLaneMs uint16

	PitStopTimerMs        uint16
	PitStopShouldServePen uint8
}

func (ld lapData) ToMessagePayload() messages.LapData {
	return messages.LapData{
		LastLapTimeMs:      int(ld.LastLapTimeMs),
		CurrentLapTimeMs:   int(ld.CurrentLapTimeMs),
		FirstSectorTimeMs:  int(ld.FirstSectorTimeMs),
		SecondSectorTimeMs: int(ld.SecondSectorTimeMs),
		LapDistance:        ld.LapDistance,
		TotalDistance:      ld.TotalDistance,
		CarPosition:        int(ld.CarPosition),
		CurrentLapNum:      int(ld.CurrentLapNum),
		Sector:             int(ld.Sector),
		DriverStatus:       int(ld.DriverStatus),
	}
}

type LapTimePacketParser struct{}

func (p LapTimePacketParser) ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error) {

	lapDataPacket := make([]lapData, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &lapDataPacket)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	lapData := lapDataPacket[header.PlayerCarIdx].ToMessagePayload()

	msg := messages.New(
		messages.LapDataMessageType,
		header.SessionUID,
		header.PacketID,
		header.FrameIdentifier,
		&lapData,
	)

	log.Printf("Lap Data: %+v\n", msg.Payload)

	return &msg, nil
}
