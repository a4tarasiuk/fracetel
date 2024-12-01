package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"

	"fracetel/core/messages"
)

type EventDataCode string

const (
	sessionStartedCode  EventDataCode = "SSTA"
	sessionFinishedCode               = "SEND"
)

type Event struct {
	Code1 uint8
	Code2 uint8
	Code3 uint8
	Code4 uint8
}

func (e Event) CodeToString() EventDataCode {
	return EventDataCode(string(e.Code1) + string(e.Code2) + string(e.Code3) + string(e.Code4))
}

func (e Event) IsSessionStarted() bool {
	return e.CodeToString() == sessionStartedCode
}

func (e Event) IsSessionFinished() bool {
	return e.CodeToString() == sessionFinishedCode
}

type eventPacketParser struct{}

func (p eventPacketParser) ToMessage(header *Header, rawPacket RawPacket) (
	*messages.Message,
	error,
) {
	event := Event{}

	// parsePacketBody(rawPacket, &event)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &event)

	if err != nil {
		log.Printf("Error during reading LapData: %s", err)
	}

	// if event.IsSessionStarted() {
	// 	log.Printf("=== SESSION STARTED ===, %s", strconv.FormatUint(header.SessionUID, 10))
	//
	// 	return &messages.Message{
	// 		Type:       messages.SessionStartedMessageType,
	// 		SessionID:  header.SessionUID,
	// 		Payload:    nil,
	// 		OccurredAt: time.Now().UTC(),
	// 	}, nil
	//
	// } else if event.IsSessionFinished() {
	// 	log.Printf("=== SESSION FINISHED ===, %s", strconv.FormatUint(header.SessionUID, 10))
	//
	// 	return &messages.Message{
	// 		Type:       messages.SessionFinishedMessageType,
	// 		SessionID:  header.SessionUID,
	// 		Payload:    nil,
	// 		OccurredAt: time.Now().UTC(),
	// 	}, nil
	// }

	return &messages.Message{}, errors.New("unsupported event")
}
