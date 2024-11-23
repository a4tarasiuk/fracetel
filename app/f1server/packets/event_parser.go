package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"strconv"
	"time"

	"fracetel/core/messages"
)

type eventPacketParser struct{}

func (p eventPacketParser) ToMessage(header *Header, rawPacket RawPacket) (
	*messages.Message,
	error,
) {
	event := Event{}

	// parsePacketBody(rawPacket, &event)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, binary.LittleEndian, &event)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	if event.IsSessionStarted() {
		log.Printf("=== SESSION STARTED ===, %s", strconv.FormatUint(header.SessionUID, 10))

		return &messages.Message{
			Type:       messages.SessionStartedMessageType,
			SessionID:  header.SessionUID,
			Payload:    nil,
			OccurredAt: time.Now().UTC(),
		}, nil

	} else if event.IsSessionFinished() {
		log.Printf("=== SESSION FINISHED ===, %s", strconv.FormatUint(header.SessionUID, 10))

		return &messages.Message{
			Type:       messages.SessionFinishedMessageType,
			SessionID:  header.SessionUID,
			Payload:    nil,
			OccurredAt: time.Now().UTC(),
		}, nil
	}

	return &messages.Message{}, errors.New("unsupported event")
}
