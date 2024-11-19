package models

type MessageType string

const (
	CarTelemetryMessageType MessageType = "CAR_TELEMETRY"
	LapDataMessageType                  = "LAP_DATA"
)
