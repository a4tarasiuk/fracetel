package models

type Message struct {
	Type MessageType `json:"type"`

	SessionID uint64 `json:"session_id"`

	Payload interface{} `json:"payload"`
}
