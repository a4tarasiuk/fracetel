package telemetry

import "time"

type Session struct {
	SessionID       uint64
	FrameIdentifier uint32
	OccurredAt      time.Time

	Weather int

	TrackTemperature int
	AirTemperature   int

	TotalLaps   int
	TrackLength int
	TrackID     int

	Type int

	TimeLeft int
	Duration int
}
