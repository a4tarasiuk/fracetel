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
	FRaceTelStreamName  = "fracetel_logs"
	FRaceTelSubjectName = FRaceTelStreamName + ".*"

	CarTelemetrySubjectName = FRaceTelStreamName + ".car_telemetry"
	LapDataSubjectName      = FRaceTelStreamName + ".lap_data"
	SessionSubjectName      = FRaceTelStreamName + ".session"
)

var MessageTypeSubjectMap = map[messages.MessageType]string{
	messages.CarTelemetryMessageType: CarTelemetrySubjectName,
	messages.LapDataMessageType:      LapDataSubjectName,
	messages.SessionMessageType:      SessionSubjectName,
}
