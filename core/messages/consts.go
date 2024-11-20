package messages

type MessageType string

const (
	CarTelemetryMessageType    MessageType = "CAR_TELEMETRY"
	LapDataMessageType                     = "LAP_DATA"
	SessionStartedMessageType              = "SESSION_STARTED"
	SessionFinishedMessageType             = "SESSION_FINISHED"
)
