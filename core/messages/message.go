package messages

import (
	"strconv"
	"time"
)

type Header struct {
	SessionID string `json:"session_id"`

	PacketID string `json:"packet_id"`

	FrameIdentifier string `json:"frame_identifier"`

	OccurredAt time.Time `json:"occurred_at"`
}

type Message struct {
	Type MessageType `json:"-"`

	Header Header `json:"header"`

	Payload interface{} `json:"payload"`
}

type genericMessage interface {
	CarTelemetry | LapData | Session | CarStatus | CarDamage | SessionHistory | FinalClassification | struct{}
}

var EmptyPayload struct{}

func New[T genericMessage](
	messageType MessageType,
	sessionID uint64,
	packetID uint8,
	frameIdentifier uint32,
	payload *T,
) Message {
	return Message{
		Type: messageType,
		Header: Header{
			SessionID:       strconv.FormatUint(sessionID, 10),
			PacketID:        strconv.Itoa(int(packetID)),
			FrameIdentifier: strconv.FormatUint(uint64(frameIdentifier), 10),
			OccurredAt:      time.Now().UTC(),
		},
		Payload: payload,
	}
}
