package packets

import (
	"errors"
	"log"

	"fracetel/core/models"
)

type eventPacketParser struct{}

func (p eventPacketParser) ToMessage(header *Header, rawPacket RawPacket) (
	*models.Message,
	error,
) {
	event := Event{}

	parsePacketBody(rawPacket, &event)

	if event.IsSessionStarted() {
		log.Printf("=== SESSION STARTED ===, %d", header.SessionUID)

		return &models.Message{
			Type:      models.SessionStartedMessageType,
			SessionID: header.SessionUID,
			Payload:   nil,
		}, nil

	} else if event.IsSessionFinished() {
		log.Printf("=== SESSION FINISHED ===, %d", header.SessionUID)

		return &models.Message{
			Type:      models.SessionFinishedMessageType,
			SessionID: header.SessionUID,
			Payload:   nil,
		}, nil
	}

	return &models.Message{}, errors.New("unsupported event")
}
