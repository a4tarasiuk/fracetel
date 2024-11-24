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

type LapData struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	LastLapTimeMs    int `bson:"last_lap_time_ms"`
	CurrentLapTimeMs int `bson:"current_lap_time_ms"`

	FirstSectorTimeMs  int `bson:"first_sector_time_ms"`
	SecondSectorTimeMs int `bson:"second_sector_time_ms"`
	Sector             int `bson:"sector"`

	LapDistance   float32 `bson:"lap_distance"`
	TotalDistance float32 `bson:"total_distance"`

	CarPosition int `bson:"car_position"`

	CurrentLapNum int `bson:"current_lap_num"`

	DriverStatus int `bson:"driver_status"`
}

func lapDataFromMessage(lapDataMessage messages.LapData, header messages.Header) LapData {
	return LapData{
		SessionID:          header.SessionID,
		FrameIdentifier:    header.FrameIdentifier,
		OccurredAt:         header.OccurredAt,
		LastLapTimeMs:      lapDataMessage.LastLapTimeMs,
		CurrentLapTimeMs:   lapDataMessage.CurrentLapTimeMs,
		FirstSectorTimeMs:  lapDataMessage.FirstSectorTimeMs,
		SecondSectorTimeMs: lapDataMessage.SecondSectorTimeMs,
		Sector:             lapDataMessage.Sector,
		LapDistance:        lapDataMessage.LapDistance,
		TotalDistance:      lapDataMessage.TotalDistance,
		CarPosition:        lapDataMessage.CarPosition,
		CurrentLapNum:      lapDataMessage.CurrentLapNum,
		DriverStatus:       lapDataMessage.DriverStatus,
	}
}

type Session struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	Weather int `bson:"weather"`

	TrackTemperature int `bson:"track_temperature"`
	AirTemperature   int `bson:"air_temperature"`

	TotalLaps   int `bson:"total_laps"`
	TrackLength int `bson:"track_length"`
	TrackID     int `bson:"track_id"`

	Type int `bson:"type"`

	TimeLeft int `bson:"time_left"`
	Duration int `bson:"duration"`
}

func sessionFromMessage(sessionMessage messages.Session, header messages.Header) Session {
	return Session{
		SessionID:        header.SessionID,
		FrameIdentifier:  header.FrameIdentifier,
		OccurredAt:       header.OccurredAt,
		Weather:          sessionMessage.Weather,
		TrackTemperature: sessionMessage.TrackTemperature,
		AirTemperature:   sessionMessage.AirTemperature,
		TotalLaps:        sessionMessage.TotalLaps,
		TrackLength:      sessionMessage.TrackLength,
		TrackID:          sessionMessage.TrackID,
		Type:             sessionMessage.Type,
		TimeLeft:         sessionMessage.TimeLeft,
		Duration:         sessionMessage.Duration,
	}
}
