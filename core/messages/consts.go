package messages

type MessageType string

const (
	CarTelemetryMessageType MessageType = "CAR_TELEMETRY"
	CarStatusMessageType    MessageType = "CAR_STATUS"
	CarDamageMessageType    MessageType = "CAR_DAMAGE"

	LapDataMessageType        MessageType = "LAP_DATA"
	SessionMessageType        MessageType = "SESSION"
	SessionHistoryMessageType MessageType = "SESSION_HISTORY"

	SessionStartedMessageType  MessageType = "SESSION_STARTED"
	SessionFinishedMessageType MessageType = "SESSION_FINISHED"
)
