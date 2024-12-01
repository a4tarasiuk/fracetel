package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"

	"fracetel/core/messages"
)

type sessionHistory struct {
	CarIdx uint8

	NumLaps uint8

	NumTyreStints uint8

	BestLapTimeLapNum uint8

	BestSector1LapNum uint8
	BestSector2LapNum uint8
	BestSector3LapNum uint8

	LapHistoryData [100]lapHistoryData

	TyreStintHistoryData [8]tyreStintHistoryData
}

func (sh sessionHistory) ToMessagePayload() messages.SessionHistory {
	lapsHistory := make([]messages.LapHistory, len(sh.LapHistoryData))

	for idx := 0; idx < len(sh.LapHistoryData); idx++ {
		lapHistoryPacket := sh.LapHistoryData[idx]

		lapsHistory[idx] = messages.LapHistory{
			LapTimeMs: int(lapHistoryPacket.LapTimeMs),
			Sector1Ms: int(lapHistoryPacket.Sector1Ms),
			Sector2Ms: int(lapHistoryPacket.Sector2Ms),
			Sector3Ms: int(lapHistoryPacket.Sector3Ms),
		}
	}

	return messages.SessionHistory{
		NumLaps:           int(sh.NumLaps),
		BestLapTimeLapNum: int(sh.BestLapTimeLapNum),
		BestSector1LapNum: int(sh.BestSector1LapNum),
		BestSector2LapNum: int(sh.BestSector2LapNum),
		BestSector3LapNum: int(sh.BestSector3LapNum),
		LapsHistory:       lapsHistory,
	}
}

type lapHistoryData struct {
	LapTimeMs uint32

	Sector1Ms uint16
	Sector2Ms uint16
	Sector3Ms uint16

	LapValidBitFlags uint8
}

type tyreStintHistoryData struct {
	EndLap uint8

	TyreActualCompound uint8
	TyreVisualCompound uint8
}

type sessionHistoryParser struct{}

func (p sessionHistoryParser) ToMessage(header *Header, rawPacket RawPacket) (*messages.Message, error) {

	sessionHistoryPacket := sessionHistory{}

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &sessionHistoryPacket)

	if err != nil {
		log.Printf("Error during reading Session: %s", err)
	}

	// Session history is sent for every car. All other cars should be ignored. Only players data must be processed
	if sessionHistoryPacket.CarIdx != header.PlayerCarIdx {
		return &messages.Message{}, errors.New("skipped as it does not relate to current player")
	}

	payload := sessionHistoryPacket.ToMessagePayload()

	msg := messages.New(
		messages.SessionHistoryMessageType,
		header.SessionUID,
		header.PacketID,
		header.FrameIdentifier,
		&payload,
	)

	log.Printf("Session history: %+v\n", msg.Payload)

	return &msg, nil
}
