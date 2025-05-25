package telemetry

import "time"

type LapData struct {
	SessionID       uint64
	FrameIdentifier uint32
	OccurredAt      time.Time

	LastLapTimeMs    int
	CurrentLapTimeMs int

	FirstSectorTimeMs  int
	SecondSectorTimeMs int

	LapDistance   float32
	TotalDistance float32

	CarPosition int

	CurrentLapNum int

	Sector int

	DriverStatus int
}
