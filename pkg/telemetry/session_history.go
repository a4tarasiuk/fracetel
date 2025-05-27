package telemetry

import "time"

type SessionHistory struct {
	SessionID       string
	FrameIdentifier uint32
	OccurredAt      time.Time

	NumLaps int

	BestLapTimeLapNum int

	BestSector1LapNum int
	BestSector2LapNum int
	BestSector3LapNum int

	LapsHistory []LapHistory
}

type LapHistory struct {
	LapTimeMs int `json:"lap_time_ms"`

	Sector1Ms int `json:"sector_1_ms"`
	Sector2Ms int `json:"sector_2_ms"`
	Sector3Ms int `json:"sector_3_ms"`
}
