package telemetry

type Session struct {
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
