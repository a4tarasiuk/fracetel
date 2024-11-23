package streams

import (
	"fracetel/core/messages"
)

const (
	SessionStreamName      = "sessions"
	SessionStreamSubject   = "sessions.*"
	SessionStartedSubject  = "sessions.started"
	SessionFinishedSubject = "sessions.finished"
)

const (
	FRaceTelStreamName      = "fracetel_logs"
	FRaceTelSubjectName     = "fracetel_logs.*"
	CarTelemetrySubjectName = "fracetel_logs.car_telemetry"
)

var MessageTypeSubjectMap = map[messages.MessageType]string{
	messages.CarTelemetryMessageType: CarTelemetrySubjectName,
}
