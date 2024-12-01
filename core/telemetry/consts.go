package telemetry

type MessageType string

const (
	CarTelemetryMessageType MessageType = "CAR_TELEMETRY"
	CarStatusMessageType    MessageType = "CAR_STATUS"
	CarDamageMessageType    MessageType = "CAR_DAMAGE"

	LapDataMessageType        MessageType = "LAP_DATA"
	SessionMessageType        MessageType = "SESSION"
	SessionHistoryMessageType MessageType = "SESSION_HISTORY"

	FinalClassificationMessageType MessageType = "FINAL_CLASSIFICATION"

	SessionStartedMessageType  MessageType = "SESSION_STARTED"
	SessionFinishedMessageType MessageType = "SESSION_FINISHED"
)

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

var MessageTypeTopicMap = map[MessageType]string{
	CarTelemetryMessageType:        CarTelemetryTopicName,
	CarStatusMessageType:           CarStatusTopicName,
	CarDamageMessageType:           CarDamageTopicName,
	LapDataMessageType:             LapDataTopicName,
	SessionMessageType:             SessionTopicName,
	SessionHistoryMessageType:      SessionHistoryTopicName,
	FinalClassificationMessageType: FinalClassificationTopicName,
}
