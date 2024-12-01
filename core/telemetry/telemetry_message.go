package telemetry

import (
	"strconv"
	"time"

	"fracetel/internal/messaging"
)

type (
	Header struct {
		SessionID string `json:"session_id"`

		FrameIdentifier string `json:"frame_identifier"`

		OccurredAt time.Time `json:"occurred_at"`
	}

	Message struct {
		Type MessageType `json:"type"`

		Header Header `json:"header"`

		Payload interface{} `json:"payload"`
	}
)

type genericPayload interface {
	CarTelemetry | LapData | Session | CarStatus | CarDamage | SessionHistory | FinalClassification | struct{}
}

func NewMessage[T genericPayload](
	messageType MessageType,
	sessionID uint64,
	frameIdentifier uint32,
	payload *T,
) Message {

	return Message{
		Type: messageType,
		Header: Header{
			SessionID:       strconv.FormatUint(sessionID, 10),
			FrameIdentifier: strconv.FormatUint(uint64(frameIdentifier), 10),
			OccurredAt:      time.Now().UTC(), // TODO: Look at packets.Header.SessionTime
		},
		Payload: payload,
	}
}

func (tm Message) GetEventName() string {
	return string(tm.Type)
}

func (tm Message) GetEventPayload() messaging.EventPayload {
	return tm
}
