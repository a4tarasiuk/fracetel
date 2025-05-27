package telemetry

import "time"

type Session struct {
	SessionID       string
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
