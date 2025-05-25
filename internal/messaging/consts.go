package messaging

import "fracetel/pkg/telemetry"

const (
	FRaceTelStreamName = "fracetel_logs"
	FRaceTelTopicName  = FRaceTelStreamName + ".*"

	CarTelemetryTopicName = FRaceTelStreamName + ".car_telemetry"
	CarStatusTopicName    = FRaceTelStreamName + ".car_status"
	CarDamageTopicName    = FRaceTelStreamName + ".car_damage"

	LapDataTopicName = FRaceTelStreamName + ".lap_data"

	SessionTopicName        = FRaceTelStreamName + ".session"
	SessionHistoryTopicName = FRaceTelStreamName + ".session_history"

	FinalClassificationTopicName = FRaceTelStreamName + ".final_classification"
)

var MessageTypeTopicMap = map[telemetry.MessageType]string{
	telemetry.CarTelemetryMessageType:        CarTelemetryTopicName,
	telemetry.CarStatusMessageType:           CarStatusTopicName,
	telemetry.CarDamageMessageType:           CarDamageTopicName,
	telemetry.LapDataMessageType:             LapDataTopicName,
	telemetry.SessionMessageType:             SessionTopicName,
	telemetry.SessionHistoryMessageType:      SessionHistoryTopicName,
	telemetry.FinalClassificationMessageType: FinalClassificationTopicName,
}
