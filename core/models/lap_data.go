package models

type LapData struct {
	LastLapTimeMs    uint32
	CurrentLapTimeMs uint32

	FirstSectorTimeMs  uint16
	SecondSectorTimeMs uint16

	LapDistance   float32
	TotalDistance float32

	CarPosition uint8

	CurrentLapNum uint8
}
