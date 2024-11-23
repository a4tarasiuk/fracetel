package worker

import (
	"time"

	"fracetel/core/messages"
)

type CarTelemetry struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	Speed int `bson:"speed"`

	Throttle float32 `bson:"throttle"`
	Steer    float32 `bson:"steer"`
	Brake    float32 `bson:"brake"`

	EngineRPM int `bson:"engine_rpm"`

	DRS byte `bson:"drs"`

	TyreSurfaceTemperature []int `bson:"tyre_surface_temperature"`

	TyreInnerTemperature []int `bson:"tyre_inner_temperature"`
}

func carTelemetryFromMessage(telemetry messages.CarTelemetry, header messages.Header) CarTelemetry {
	carTelemetry := CarTelemetry{
		SessionID:              header.SessionID,
		FrameIdentifier:        header.FrameIdentifier,
		OccurredAt:             header.OccurredAt,
		Speed:                  telemetry.Speed,
		Throttle:               telemetry.Throttle,
		Steer:                  telemetry.Steer,
		Brake:                  telemetry.Brake,
		EngineRPM:              telemetry.EngineRPM,
		DRS:                    telemetry.DRS,
		TyreSurfaceTemperature: telemetry.TyreSurfaceTemperature,
		TyreInnerTemperature:   telemetry.TyreInnerTemperature,
	}

	return carTelemetry
}
