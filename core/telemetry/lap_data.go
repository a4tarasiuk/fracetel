package telemetry

type LapData struct {
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
