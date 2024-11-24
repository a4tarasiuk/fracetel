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
	CarStatusSubjectName    = FRaceTelStreamName + ".car_status"
	CarDamageSubjectName    = FRaceTelStreamName + ".car_damage"

	LapDataSubjectName = FRaceTelStreamName + ".lap_data"

	SessionSubjectName        = FRaceTelStreamName + ".session"
	SessionHistorySubjectName = FRaceTelStreamName + ".session_history"
)

var MessageTypeSubjectMap = map[messages.MessageType]string{
	messages.CarTelemetryMessageType:   CarTelemetrySubjectName,
	messages.CarStatusMessageType:      CarStatusSubjectName,
	messages.CarDamageMessageType:      CarDamageSubjectName,
	messages.LapDataMessageType:        LapDataSubjectName,
	messages.SessionMessageType:        SessionSubjectName,
	messages.SessionHistoryMessageType: SessionHistorySubjectName,
}
