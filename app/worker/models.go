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

type CarStatus struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	TractionControl int `bson:"traction_control"`
	AntiLockBrakes  int `bson:"anti_lock_brakes"`

	FuelMix int `bson:"fuel_mix"`

	FrontBrakeBias int `bson:"front_brake_bias"`

	PutLimiterStatus int `bson:"put_limiter_status"`

	FuelInTank        float32 `bson:"fuel_in_tank"`
	FuelCapacity      float32 `bson:"fuel_capacity"`
	FuelRemainingLaps float32 `bson:"fuel_remaining_laps"`

	MaxRPM  int `bson:"max_rpm"`
	IdleRPM int `bson:"idle_rpm"`

	MaxGears int `bson:"max_gears"`

	DRSAllowed            int `bson:"drs_allowed"`
	DRSActivationDistance int `bson:"drs_activation_distance"`

	ActualTyreCompound int `bson:"actual_tyre_compound"`
	VisualTyreCompound int `bson:"visual_tyre_compound"`

	TyresAgeLaps int `bson:"tyres_age_laps"`

	VehicleFIAFlags int `bson:"vehicle_fia_flags"`

	ERSStoreEnergy float32 `bson:"ers_store_energy"`
	ERSDeployMode  int     `bson:"ers_deploy_mode"`

	ERSHarvestedThisLapMGUK float32 `bson:"ers_harvested_this_lap_mguk"`
	ERSHarvestedThisLapMGUH float32 `bson:"ers_harvested_this_lap_mguh"`
	ERSDeployedThisLap      float32 `bson:"ers_deployed_this_lap"`
}

func carStatusFromMessage(carStatusMessage messages.CarStatus, header messages.Header) CarStatus {
	return CarStatus{
		SessionID:               header.SessionID,
		FrameIdentifier:         header.FrameIdentifier,
		OccurredAt:              header.OccurredAt,
		TractionControl:         carStatusMessage.TractionControl,
		AntiLockBrakes:          carStatusMessage.AntiLockBrakes,
		FuelMix:                 carStatusMessage.FuelMix,
		FrontBrakeBias:          carStatusMessage.FrontBrakeBias,
		PutLimiterStatus:        carStatusMessage.PutLimiterStatus,
		FuelInTank:              carStatusMessage.FuelInTank,
		FuelCapacity:            carStatusMessage.FuelCapacity,
		FuelRemainingLaps:       carStatusMessage.FuelRemainingLaps,
		MaxRPM:                  carStatusMessage.MaxRPM,
		IdleRPM:                 carStatusMessage.IdleRPM,
		MaxGears:                carStatusMessage.MaxGears,
		DRSAllowed:              carStatusMessage.DRSAllowed,
		DRSActivationDistance:   carStatusMessage.DRSActivationDistance,
		ActualTyreCompound:      carStatusMessage.ActualTyreCompound,
		VisualTyreCompound:      carStatusMessage.VisualTyreCompound,
		TyresAgeLaps:            carStatusMessage.TyresAgeLaps,
		VehicleFIAFlags:         carStatusMessage.VehicleFIAFlags,
		ERSStoreEnergy:          carStatusMessage.ERSStoreEnergy,
		ERSDeployMode:           carStatusMessage.ERSDeployMode,
		ERSHarvestedThisLapMGUK: carStatusMessage.ERSHarvestedThisLapMGUK,
		ERSHarvestedThisLapMGUH: carStatusMessage.ERSHarvestedThisLapMGUH,
		ERSDeployedThisLap:      carStatusMessage.ERSDeployedThisLap,
	}
}

type CarDamage struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	TyresWear []int `bson:"tyres_wear"`

	Tyres []int `bson:"tyres"`

	Brakes []int `bson:"brakes"`
}

func carDamageFromMessage(carDamageMessage messages.CarDamage, header messages.Header) CarDamage {
	return CarDamage{
		SessionID:       header.SessionID,
		FrameIdentifier: header.FrameIdentifier,
		OccurredAt:      header.OccurredAt,
		TyresWear:       carDamageMessage.TyresWear,
		Tyres:           carDamageMessage.Tyres,
		Brakes:          carDamageMessage.Brakes,
	}
}

type SessionHistory struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	NumLaps int `bson:"num_laps"`

	BestLapTimeLapNum int `bson:"best_lap_time_lap_num"`

	BestSector1LapNum int `bson:"best_sector_1_lap_num"`
	BestSector2LapNum int `bson:"best_sector_2_lap_num"`
	BestSector3LapNum int `bson:"best_sector_3_lap_num"`

	LapsHistory []LapHistory `bson:"laps_history"`
}

type LapHistory struct {
	LapTimeMs int `bson:"lap_time_ms"`

	Sector1Ms int `bson:"sector_1_ms"`
	Sector2Ms int `bson:"sector_2_ms"`
	Sector3Ms int `bson:"sector_3_ms"`
}

func sessionHistoryFromMessage(sessionHistoryMessage messages.SessionHistory, header messages.Header) SessionHistory {
	lapsHistory := make([]LapHistory, len(sessionHistoryMessage.LapsHistory))

	for idx := 0; idx < len(sessionHistoryMessage.LapsHistory); idx++ {
		lapHistoryPacket := sessionHistoryMessage.LapsHistory[idx]

		lapsHistory[idx] = LapHistory{
			LapTimeMs: lapHistoryPacket.LapTimeMs,
			Sector1Ms: lapHistoryPacket.Sector1Ms,
			Sector2Ms: lapHistoryPacket.Sector2Ms,
			Sector3Ms: lapHistoryPacket.Sector3Ms,
		}
	}

	return SessionHistory{
		SessionID:         header.SessionID,
		FrameIdentifier:   header.FrameIdentifier,
		OccurredAt:        header.OccurredAt,
		NumLaps:           sessionHistoryMessage.NumLaps,
		BestLapTimeLapNum: sessionHistoryMessage.BestLapTimeLapNum,
		BestSector1LapNum: sessionHistoryMessage.BestSector1LapNum,
		BestSector2LapNum: sessionHistoryMessage.BestSector2LapNum,
		BestSector3LapNum: sessionHistoryMessage.BestSector3LapNum,
		LapsHistory:       lapsHistory,
	}
}

type FinalClassification struct {
	SessionID       string    `bson:"session_id"`
	FrameIdentifier string    `bson:"frame_identifier"`
	OccurredAt      time.Time `bson:"occurred_at"`

	FinishingPosition int `bson:"finishing_position"`

	StartingPosition int `bson:"starting_position"`

	BestLapTimeMs float32 `bson:"best_lap_time_ms"`
}

func finalClassificationFromMessage(
	finalClassificationMessage messages.FinalClassification,
	header messages.Header,
) FinalClassification {
	return FinalClassification{
		SessionID:         header.SessionID,
		FrameIdentifier:   header.FrameIdentifier,
		OccurredAt:        header.OccurredAt,
		FinishingPosition: finalClassificationMessage.FinishingPosition,
		StartingPosition:  finalClassificationMessage.StartingPosition,
		BestLapTimeMs:     finalClassificationMessage.BestLapTimeMs,
	}
}
