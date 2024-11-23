package messages

import "time"

type Message struct {
	Type MessageType `json:"type"`

	SessionID uint64 `json:"session_id"`

	Payload interface{} `json:"payload"`

	OccurredAt time.Time `json:"occurred_at"`
}

type genericMessage interface {
	CarTelemetry | LapData | struct{}
}

var EmptyPayload struct{}

func New[T genericMessage](messageType MessageType, sessionID uint64, payload *T) Message {
	return Message{
		Type:       messageType,
		SessionID:  sessionID,
		Payload:    payload,
		OccurredAt: time.Now().UTC(),
	}
}
