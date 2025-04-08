package telemetry

type SessionHistory struct {
	NumLaps int

	BestLapTimeLapNum int

	BestSector1LapNum int
	BestSector2LapNum int
	BestSector3LapNum int

	LapsHistory []LapHistory
}

type LapHistory struct {
	LapTimeMs int

	Sector1Ms int
	Sector2Ms int
	Sector3Ms int
}
